package gouse

import (
	"context"
	"fmt"
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

type IBackend struct {
	URL    string `json:"url"`
	IsDead bool
	mutex  sync.RWMutex
}

type ILbConfig struct {
	ProxyPort string     `json:"proxy"`
	Backends  []IBackend `json:"backends"`
}

func (backend *IBackend) SetDead(b bool) {
	backend.mutex.Lock()
	backend.IsDead = b
	backend.mutex.Unlock()
}

func (backend *IBackend) GetIsDead() bool {
	backend.mutex.RLock()
	isAlive := backend.IsDead
	backend.mutex.RUnlock()
	return isAlive
}

var mutex sync.Mutex
var idx int = 0
var cfg ILbConfig

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

func LoadBalancer(proxyPort string, backends []IBackend) {
	cfg = ILbConfig{
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
		time.Sleep(igs.SleepTimout) // 5 * time.Second
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	fmt.Println(igs.StartMsg)

	srv := http.Server{
		Addr:              ":" + igs.Port,
		ReadHeaderTimeout: 5 * time.Second, // mitigate Slowloris attack by set timeout
	}

	go func() {
		<-c
		fmt.Println(igs.ShutdownMsg)
		go func() {
			for {
				fmt.Println("waiting for goroutines to finish...")
				time.Sleep(igs.HeaderTimeout) // 1 * time.Second
			}
		}()
		wg.Wait()

		if err := srv.Shutdown(context.Background()); err != nil {
			fmt.Println("server shutdown error:", err)
		}
	}()
	fmt.Println(srv.ListenAndServe())
}
