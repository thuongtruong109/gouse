package gouse

import (
	"context"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"time"
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

func IsAlive(url *url.URL) bool {
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
			isAlive := IsAlive(pingURL)
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
	signal.Notify(c, os.Interrupt)

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
