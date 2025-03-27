package samples

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
)

/*
Description: Read JSON configuration file
Input params: (path string, v interface{})
*/
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

/*
Description: Read TOML configuration file
Input params: (path string, v interface{})
*/
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

/*
Description: Read YAML configuration file
Input params: (path string, v interface{})
*/
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
