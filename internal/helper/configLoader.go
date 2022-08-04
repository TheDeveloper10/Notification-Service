package helper

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	Server struct {
		Addr string
	}
	Database struct {
		Driver string 
		Username string 
		Password string 
		Host string 
		Name string 
		PoolSize int `yaml:"pool_size"`
	}
}

var Config config

func LoadConfig(fileName string) bool {
	file, err := os.Open(fileName)
	if err != nil {
		return false
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return false
	}

	Config = config{}
	yaml.Unmarshal(data, &Config)

	return true
}