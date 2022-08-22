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
		UseSMS      string `yaml:"use_sms"`
	}
	HTTPServer struct {
		Addr                  string
		AccessTokenExpiryTime int `yaml:"access_token_expiry_time"`
	} `yaml:"http_server"`
	RabbitMQ struct {
		URL                   string
		TemplatesQueueMax     int `yaml:"templates_queue_max"`
		NotificationsQueueMax int `yaml:"notifications_queue_max"`
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
	Twillio struct {
		MessagingServiceSID string `yaml:"messaging_service_sid"`
		AccountSID          string `yaml:"account_sid"`
		AuthToken           string `yaml:"authentication_token"`
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
