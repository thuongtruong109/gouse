
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Server' />


```go
import (
	"time"
	"github.com/thuongtruong109/gouse"
)
```

## 1. Server loadbalancer

Description: Load balancer with health check<br>Input params: (ILbConf.ProxyPort, ILbConf.Backends)<br>

```go
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
```

## 2. Graceful shutdown server

Description: Start server with graceful shutdown mode for API<br>Input params: (gouse.IServer)<br>

```go
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
```

## 3. Server proxy

Description: Serve proxy wrapper forward to another port<br>Input params: (port string, urls []string{})<br>

```go
func ServerProxy() {
	gouse.Proxy("5000", []string{"http://localhost:3000", "http://localhost:3001"})
}
```
