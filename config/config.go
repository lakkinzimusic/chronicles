package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
		Protocol string `yaml:"protocol"`
	} `yaml:"server"`
	DB struct {
		Database string `yaml:"database"`
		Collection string `yaml:"collection"`
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"DB"`
}

func getConfigFile(name string) []byte{
	pwd, _ := os.Getwd()
	yamlFile, err := ioutil.ReadFile(filepath.Join(pwd, name))
	if err != nil {
		panic(err)
	}
	return yamlFile
}

func GetConfig(fileName string ) Config{
	file := getConfigFile(fileName)
	var cfg Config
	err := yaml.Unmarshal(file, &cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}