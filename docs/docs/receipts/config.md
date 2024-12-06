
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Config' />


```go
import (
	"fmt"
	"github.com/thuongtruong109/gouse"
)
```

## 1. Config json

Description: Read JSON configuration file<br>

```go
func ConfigJson() {
	type configuration struct {
		Users  []string
		Groups []string
	}

	var myConf configuration
	err := gouse.ReadJSON("conf.json", &myConf)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Users:", myConf.Users)
	fmt.Println("Groups:", myConf.Groups)
}
```

## 2. Config toml

Description: Read TOML configuration file<br>

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
		fmt.Println("Error:", err)
	}

	fmt.Println("Mysql Host:", myConf.Mysql.Host)
	fmt.Println("Mysql Username:", myConf.Mysql.Username)
	fmt.Println("Mysql Password:", myConf.Mysql.Password)
	fmt.Println("Mysql Database:", myConf.Mysql.Database)
}
```

## 3. Config yaml

Description: Read YAML configuration file<br>

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
		fmt.Println("Error:", err)
	}

	fmt.Println("Server Port:", myConf.Server.Port)
	fmt.Println("Server Host:", myConf.Server.Host)
	fmt.Println("DB Username:", myConf.Database.Username)
	fmt.Println("DB Password:", myConf.Database.Password)
}
```
