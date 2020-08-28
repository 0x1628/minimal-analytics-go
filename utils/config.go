package utils

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type ConfigItem struct {
	AccessKeyID     string `yaml:"AccessKeyID"`
	AccessKeySecret string `yaml:"AccessKeySecret"`
	Endpoint        string `yaml:"Endpoint"`
	ProjectName     string `yaml:"ProjectName"`
	LogstoreName    string `yaml:"LogstoreName"`
}

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	Puda ConfigItem `yaml:"puda"`
}

var config Config

func InitConfig() Config {
	f, err := os.Open("./config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var c Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	config = c
	return config
}

func GetConfig() Config {
	return config
}
