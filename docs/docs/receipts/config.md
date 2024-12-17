
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Config' />


```go
import (
	"github.com/thuongtruong109/gouse"
)
```

## 1. Config json

Description: Read JSON configuration file<br>Input params: (path string, v interface{})<br>

```go
func ConfigJson() {
	type configuration struct {
		Users  []string
		Groups []string
	}

	var myConf configuration
	err := gouse.ReadJSON("conf.json", &myConf)
	if err != nil {
		gouse.Println("Error:", err)
	}

	gouse.Println("Users:", myConf.Users)
	gouse.Println("Groups:", myConf.Groups)
}
```

## 2. Config toml

Description: Read TOML configuration file<br>Input params: (path string, v interface{})<br>

```go
func ConfigToml() {
	type configuration struct {
		Mysql struct {
			Host     string
			Username string
			Password string
			Database string
		}
	}

	var myConf configuration
	err := gouse.ReadTOML("conf.toml", &myConf)
	if err != nil {
		gouse.Println("Error:", err)
	}

	gouse.Println("Mysql Host:", myConf.Mysql.Host)
	gouse.Println("Mysql Username:", myConf.Mysql.Username)
	gouse.Println("Mysql Password:", myConf.Mysql.Password)
	gouse.Println("Mysql Database:", myConf.Mysql.Database)
}
```

## 3. Config yaml

Description: Read YAML configuration file<br>Input params: (path string, v interface{})<br>

```go
func ConfigYaml() {
	type configuration struct {
		Server struct {
			Port string `envconfig:"SERVER_PORT"`
			Host string `envconfig:"SERVER_HOST"`
		}
		Database struct {
			Username string `envconfig:"DATABASE_USERNAME"`
			Password string `envconfig:"DATABASE_PASSWORD"`
		}
	}

	var myConf configuration
	err := gouse.ReadYAML("conf.yaml", &myConf)
	if err != nil {
		gouse.Println("Error:", err)
	}

	gouse.Println("Server Port:", myConf.Server.Port)
	gouse.Println("Server Host:", myConf.Server.Host)
	gouse.Println("DB Username:", myConf.Database.Username)
	gouse.Println("DB Password:", myConf.Database.Password)
}
```
