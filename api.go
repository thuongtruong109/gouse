package gouse

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ILb struct {
	URL    string `json:"url"`
	IsDead bool
	mutex  sync.RWMutex
}

type ILbConf struct {
	ProxyPort string `json:"proxy"`
	Backends  []ILb  `json:"backends"`
}

func (backend *ILb) SetDead(b bool) {
	backend.mutex.Lock()
	backend.IsDead = b
	backend.mutex.Unlock()
}

func (backend *ILb) GetIsDead() bool {
	backend.mutex.RLock()
	isAlive := backend.IsDead
	backend.mutex.RUnlock()
	return isAlive
}

var mutex sync.Mutex
var idx int = 0
var cfg ILbConf

func _lbHandler(w http.ResponseWriter, r *http.Request) {
	maxLen := len(cfg.Backends)

	// Round Robin
	mutex.Lock()
	currentBackend := &cfg.Backends[idx%maxLen]
	if currentBackend.GetIsDead() {
		idx++
	}

	targetURL, err := url.Parse(currentBackend.URL)
	if err != nil {
		log.Println(err.Error())
	}
	idx++
	mutex.Unlock()
	reverseProxy := httputil.NewSingleHostReverseProxy(targetURL)
	reverseProxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, e error) {
		log.Printf("%v is dead", targetURL)
		currentBackend.SetDead(true)
		_lbHandler(w, r)
	}
	reverseProxy.ServeHTTP(w, r)
}

func _isAlive(url *url.URL) bool {
	conn, err := net.DialTimeout("tcp", url.Host, time.Minute*1)
	if err != nil {
		log.Printf("Unreachable tp %v, error:%v", url.Host, err.Error())
		return false
	}

	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()

	return true
}

func HealthCheck() {
	t := time.NewTicker(time.Minute * 1)

	for range t.C {
		for i := range cfg.Backends {
			backend := &cfg.Backends[i]
			pingURL, err := url.Parse(backend.URL)
			if err != nil {
				log.Print(err.Error())
				continue
			}
			isAlive := _isAlive(pingURL)
			backend.SetDead(!isAlive)
			msg := "alive"
			if !isAlive {
				msg = "dead"
			}
			log.Printf("%v checked %v by healthcheck", backend.URL, msg)
		}
	}
}

func LoadBalancer(proxyPort string, backends []ILb) {
	cfg = ILbConf{
		ProxyPort: proxyPort,
		Backends:  backends,
	}

	go HealthCheck()

	s := http.Server{
		Addr:              ":" + cfg.ProxyPort,
		Handler:           http.HandlerFunc(_lbHandler),
		ReadHeaderTimeout: 10 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Printf("Error starting server: %s", err.Error())
	}
}

type IGracefulShutdown struct {
	Port          string
	StartMsg      string
	ShutdownMsg   string
	SleepTimout   time.Duration
	HeaderTimeout time.Duration
}

func (igs *IGracefulShutdown) GracefulShutdown() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		time.Sleep(igs.SleepTimout)
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	Println(igs.StartMsg)

	srv := http.Server{
		Addr:              ":" + igs.Port,
		ReadHeaderTimeout: igs.HeaderTimeout,
	}

	go func() {
		<-c
		Println(igs.ShutdownMsg)
		go func() {
			for {
				Println("waiting for goroutines to finish...")
				time.Sleep(1 * time.Second)
			}
		}()
		wg.Wait()

		if err := srv.Shutdown(context.Background()); err != nil {
			Println("server shutdown error:", err)
		}
	}()
	Println(srv.ListenAndServe())
}

func Validate(ctxBind func() error, ctxReq func() context.Context, req any) error {
	validate := validator.New()

	if err := ctxBind(); err != nil {
		return err
	}

	ctx := ctxReq()

	return validate.StructCtx(ctx, req)
}

func Proxy(port string, urls []string) {
	_handleRequest := func(urls []string) gin.HandlerFunc {
		var counter uint64

		return func(ctx *gin.Context) {
			path := ctx.Param("path")
			if path == "" {
				ctx.IndentedJSON(http.StatusBadRequest, gin.H{
					"message": "Path is required",
				})
				ctx.Done()
				return
			}

			index := atomic.AddUint64(&counter, 1) % uint64(len(urls))
			requestedURL := urls[index] + path[1:]

			Println("Requested URL: ", requestedURL)

			req, _ := http.NewRequest(ctx.Request.Method, requestedURL, ctx.Request.Body)

			req.Header = ctx.Request.Header.Clone()
			req.Header.Del("origin")
			req.Header.Del("referer")

			queryValues := req.URL.Query()
			for k, v := range ctx.Request.URL.Query() {
				queryValues.Add(k, v[0])
			}
			req.URL.RawQuery = queryValues.Encode()

			response, err1 := http.DefaultClient.Do(req)

			for k, v := range response.Header.Clone() {
				ctx.Header(k, v[0])
			}

			ctx.Header("Access-Control-Allow-Origin", "*")
			ctx.Header("Access-Control-Allow-Methods", "*")
			ctx.Header("Access-Control-Allow-Headers", "*")

			responseBytes, err2 := io.ReadAll(response.Body)

			if err1 != nil || err2 != nil {
				ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
					"message": "Failed to request",
				})
				ctx.Done()
				return
			}

			ctx.Data(response.StatusCode, response.Header.Get("Content-Type"), responseBytes)
		}
	}

	router := gin.Default()

	router.GET("*path", _handleRequest(urls))
	router.POST("*path", _handleRequest(urls))
	router.PUT("*path", _handleRequest(urls))
	router.PATCH("*path", _handleRequest(urls))
	router.DELETE("*path", _handleRequest(urls))
	router.OPTIONS("*path", _handleRequest(urls))
	router.HEAD("*path", _handleRequest(urls))

	router.Run(":" + port)
}

func _httpError(w http.ResponseWriter, msg string, statusCode int) {
	http.Error(w, msg, statusCode)
}

func _generateFileName(originalFileName string) string {
	fileExt := filepath.Ext(originalFileName)
	originalName := strings.TrimSuffix(filepath.Base(originalFileName), fileExt)
	now := time.Now().UnixNano()
	filename := fmt.Sprintf("%s-%d%s", strings.ReplaceAll(strings.ToLower(originalName), " ", "-"), now, fileExt)
	return filename
}

func UploadSingle(w http.ResponseWriter, r *http.Request, servePath, storePath string) {
	err := r.ParseMultipartForm(10 << 20) // Limit to 10MB
	if err != nil {
		_httpError(w, fmt.Sprintf("Error parsing form: %s", err.Error()), http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		_httpError(w, fmt.Sprintf("Error retrieving file: %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer file.Close()

	filename := _generateFileName(header.Filename)
	filePath := filepath.Join(servePath, filename)

	out, err := os.Create(filepath.Join(storePath, filename))
	if err != nil {
		_httpError(w, fmt.Sprintf("Unable to create file: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		_httpError(w, fmt.Sprintf("Error writing file: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"filepath": "%s"}`, filePath)))
}

func UploadMulti(w http.ResponseWriter, r *http.Request, servePath, storePath string) {
	err := r.ParseMultipartForm(10 << 20) // Limit to 10MB
	if err != nil {
		_httpError(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	files := r.MultipartForm.File["images"]
	if len(files) == 0 {
		_httpError(w, "No files uploaded", http.StatusBadRequest)
		return
	}

	var filePaths []string

	for _, file := range files {
		filename := _generateFileName(file.Filename)
		filePath := filepath.Join(servePath, filename)
		filePaths = append(filePaths, filePath)

		out, err := os.Create(filepath.Join(storePath, filename))
		if err != nil {
			_httpError(w, fmt.Sprintf("Error saving file: %s", err.Error()), http.StatusInternalServerError)
			return
		}
		defer out.Close()

		readerFile, err := file.Open()
		if err != nil {
			_httpError(w, fmt.Sprintf("Error opening file: %s", err.Error()), http.StatusInternalServerError)
			return
		}
		defer readerFile.Close()

		_, err = io.Copy(out, readerFile)
		if err != nil {
			_httpError(w, fmt.Sprintf("Error copying file content: %s", err.Error()), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"filepath": %v}`, filePaths)))
}
