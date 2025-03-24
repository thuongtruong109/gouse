package samples

import (
	"net/http"
	"time"

	"github.com/thuongtruong109/gouse"
)

/*
Description: Load balancer with health check
Input params: (ILbConf.ProxyPort, ILbConf.Backends)
*/
func Loadbalancer() {
	lbCfg := gouse.ILbConf{
		ProxyPort: "8080",
		Backends: []gouse.ILb{
			{
				URL:    "http://localhost:8081",
				IsDead: false,
			},
			{
				URL:    "http://localhost:8082",
				IsDead: false,
			},
			{
				URL:    "http://localhost:8083",
				IsDead: false,
			},
		},
	}

	// gouse.HealthCheck() // Enable health check (optional)
	gouse.LoadBalancer(lbCfg.ProxyPort, lbCfg.Backends)
}

/*
Description: Graceful shutdown for API server
Input params: (IGracefulShutdown)
*/
func GracefulShutdown() {
	gs := gouse.IGracefulShutdown{
		Port:          "3000",
		StartMsg:      "Starting server at port http://localhost:3000",
		ShutdownMsg:   "Shutting down server...",
		SleepTimout:   5 * time.Second,
		HeaderTimeout: 5 * time.Second,
	}

	gs.GracefulShutdown()
}

/*
Description: Upload file to server
Input params: (servePath, storagePath string)
*/
func UploadFile() {
	{
		http.HandleFunc("/upload-single", func(w http.ResponseWriter, r *http.Request) {
			gouse.UploadSingle(w, r, "http://localhost:8080/img", "./public")
		})
		http.HandleFunc("/upload-multi", func(w http.ResponseWriter, r *http.Request) {
			gouse.UploadMulti(w, r, "http://localhost:8080/img", "./public")
		})

		http.ListenAndServe(":8080", nil)
	}
}
