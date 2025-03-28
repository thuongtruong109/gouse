package samples

import (
	"time"

	"github.com/thuongtruong109/gouse"
)

/*
Description: Load balancer with health check
Input params: (ILbConf.ProxyPort, ILbConf.Backends)
*/
func ServerLoadbalancer() {
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
Description: Start server with graceful shutdown mode for API
Input params: (gouse.IServer)
*/
func GracefulShutdownServer() {
	gs := gouse.IServer{
		Port:          "3000",
		StartMsg:      "Starting server at port http://localhost:3000",
		ShutdownMsg:   "Shutting down server...",
		SleepTimeout:  5 * time.Second,
		HeaderTimeout: 5 * time.Second,
	}

	gs.Server()
}

/*
Description: Serve proxy wrapper forward to another port
Input params: (port string, urls []string{})
*/
func ServerProxy() {
	gouse.Proxy("5000", []string{"http://localhost:3000", "http://localhost:3001"})
}
