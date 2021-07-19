package util

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type SlsConfig struct {
	AccessKeyID      string `yaml:"AccessKeyID"`
	AccessKeySecret  string `yaml:"AccessKeySecret"`
	Endpoint         string `yaml:"Endpoint"`
	ProjectName      string `yaml:"ProjectName"`
	EventLogstore    string `yaml:"EventLogstore"`
	RegisterLogstore string `yaml:"RegisterLogstore"`
	CrashLogstore    string `yaml:"CrashLogstore"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Config struct {
	Server ServerConfig `yaml:"server"`
	Sls    SlsConfig    `yaml:"sls"`
}

var config Config

func init() {
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
}

func GetConfig() Config {
	return config
}
