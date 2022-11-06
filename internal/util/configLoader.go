package util

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	Service struct {
		Services          	 ListOfStrings `yaml:"services"`
		Clients           	 ListOfStrings `yaml:"clients"`
		AllowedLanguages  	 ListOfStrings `yaml:"allowed_languages"`
		TemplateCacheSize 	 int           `yaml:"template_cache_size"`

		AccessTokenSecret 	 string        `yaml:"access_token_secret"`
		AccessTokenExpiry 	 int           `yaml:"access_token_expiry"`
		RefresherTokenSecret string        `yaml:"refresher_token_secret"`
		RefresherTokenExpiry int           `yaml:"refresher_token_expiry"`
	}
	HTTPServer struct {
		Addr                  string
		MasterAccessToken     string `yaml:"master_access_token"`
		AccessTokenExpiryTime int    `yaml:"access_token_expiry_time"`
	} `yaml:"http_server"`
	RabbitMQ struct {
		URL                    string
		TemplatesQueueName     string `yaml:"templates_queue_name"`
		TemplatesQueueMax      int 	  `yaml:"templates_queue_max"`
		NotificationsQueueName string `yaml:"notifications_queue_name"`
		NotificationsQueueMax  int 	  `yaml:"notifications_queue_max"`
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
		logrus.Fatal(err.Error())
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	Config = config{}
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		logrus.Fatal(err.Error())
	}
}

const ServiceConfigPath = "./config/service_config.yaml"