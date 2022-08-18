package helper

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	Service struct {
		UseHTTP     string `yaml:"use_http"`
		UseRabbitMQ string `yaml:"use_rabbitmq"`
		UseSMTP     string `yaml:"use_smtp"`
		UsePush     string `yaml:"use_push"`
	}
	HTTPServer struct {
		Addr string
	} `yaml:"http_server"`
	RabbitMQ struct {
		URL string
	}
	Database struct {
		Driver   string
		Username string
		Password string
		Host     string
		Name     string
		PoolSize int `yaml:"pool_size"`
	}
	SMTP struct {
		FromEmail    string `yaml:"from_email"`
		FromPassword string `yaml:"from_password"`
		Host         string
		Port         int
	}
}

var Config config

func LoadConfig(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err.Error())
	}

	Config = config{}
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		log.Fatal(err.Error())
	}
}
