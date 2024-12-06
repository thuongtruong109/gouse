
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Api' />


```go
import (
	"time"	"github.com/thuongtruong109/gouse")
```

## 1. Api loadbalancer

Description: Load balancer with health check<br>Input params: (ILbConfig.ProxyPort, ILbConfig.Backends)<br>

```go
func ApiLoadbalancer() {
	lbCfg := gouse.ILbConfig{
		ProxyPort: "8080",
		Backends: []gouse.IBackend{
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
}```

## 2. Api graceful shutdown

Description: Graceful shutdown for API server<br>Input params: (IGracefulShutdown)<br>

```go
func ApiGracefulShutdown() {
	gs := gouse.IGracefulShutdown{
		Port:          "3000",
		StartMsg:      "Starting server at port http://localhost:3000",
		ShutdownMsg:   "Shutting down server...",
		SleepTimout:   5 * time.Second,
		HeaderTimeout: 5 * time.Second,
	}

	gs.GracefulShutdown()
}```
