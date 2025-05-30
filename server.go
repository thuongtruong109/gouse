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
	"runtime"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
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

func lbHandler(w http.ResponseWriter, r *http.Request) {
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
		lbHandler(w, r)
	}
	reverseProxy.ServeHTTP(w, r)
}

func isAlive(url *url.URL) bool {
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
			isAlive := isAlive(pingURL)
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
		Handler:           http.HandlerFunc(lbHandler),
		ReadHeaderTimeout: 10 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Printf("Error starting server: %s", err.Error())
	}
}

type IServer struct {
	Port          string
	StartMsg      string
	ShutdownMsg   string
	SleepTimeout  time.Duration
	HeaderTimeout time.Duration
}

func (igs *IServer) Server() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		time.Sleep(igs.SleepTimeout)
	}()

	debug.FreeOSMemory()
	runtime.GC()
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGHUP, syscall.SIGQUIT)

	fmt.Println(igs.StartMsg)

	srv := http.Server{
		Addr:              ":" + igs.Port,
		ReadHeaderTimeout: igs.HeaderTimeout,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(igs.SleepTimeout))
	defer cancel()

	p, _ := os.FindProcess(os.Getpid())
	p.Signal(syscall.SIGINT)

	go func() {
		<-c
		fmt.Println(igs.ShutdownMsg)
		go func() {
			for {
				fmt.Println("waiting for goroutines to finish...")
				time.Sleep(1 * time.Second)
			}
		}()
		wg.Wait()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("server shutdown error: %v", fmt.Sprintf("%v", err))
		}
	}()
	fmt.Println(srv.ListenAndServe())
}

func Proxy(port string, urls []string) {
	handleRequest := func(urls []string) gin.HandlerFunc {
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

			fmt.Println("Requested URL: ", requestedURL)

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

	router.GET("*path", handleRequest(urls))
	router.POST("*path", handleRequest(urls))
	router.PUT("*path", handleRequest(urls))
	router.PATCH("*path", handleRequest(urls))
	router.DELETE("*path", handleRequest(urls))
	router.OPTIONS("*path", handleRequest(urls))
	router.HEAD("*path", handleRequest(urls))

	router.Run(":" + port)
}
