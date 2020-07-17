package config

import (
	"github.com/jinzhu/configor"
)

var Config = struct {
	HTTPS    bool   `default:"false" env:"HTTPS"` // 是否为HTTPS
	Certpath string `default:"" env:"Certpath"`
	Certkey  string `default:"" env:"Certkey"`
	Port     uint   `default:"8080" env:"PORT"` // 启动的端口
	Host     string `default:"" env:"Host"`     // 启动的HOST
	DB       struct {
		Name     string `env:"DBName" default:"iris"`
		Adapter  string `env:"DBAdapter" default:"mysql"`
		Host     string `env:"DBHost" default:"localhost"`
		Port     string `env:"DBPort" default:"3306"`
		User     string `env:"DBUser" default:"root"`
		Password string `env:"DBPassword" default:"root"`
		Prefix   string `default:""`
	}
}{}

func init() {
	configPath := "./config/application.yml"
	if err := configor.Load(&Config, configPath); err != nil {
		panic(err)
	}
	//fmt.Println(Config)
}
