package configs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

type Server struct {
	Timeout time.Duration `yaml:"timeout"`
	Port    string        `yaml:"port"`
}

type Service struct {
	Url    string `yaml:"url"`
	Method string `yaml:"method"`
}

type AccessExceptions struct {
	List []string `yaml:"list"`
}

type DataBase struct {
	HostDb   string `yaml:"host_db"`
	PortDb   string `yaml:"port_db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	SslMode  string `yaml:"sslmode"`
}

type Config struct {
	Server           `yaml:"server"`
	Handlers         map[string]Service `yaml:"handlers"`
	Api              map[string]Service `yaml:"api"`
	AccessExceptions `yaml:"accessExceptions"`
	DataBase         `yaml:"db"`
}

var config Config

func Get() *Config {
	return &config
}

func init() {
	file, err := ioutil.ReadFile("./internal/configs/config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}
	log.Println("Config is loaded")
}
