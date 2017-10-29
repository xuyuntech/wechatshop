package config

import (
	"io/ioutil"
	"fmt"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Addr string `yaml:"addr,omitempty"`
		Debug bool `yaml:"debug"`
	}    `yaml:"server,omitempty"`
	DB struct{
		Host string `yaml:"host,omitempty"`
		Port int `yaml:"port,omitempty"`
		User string `yaml:"user,omitempty"`
		Password string `yaml:"password,omitempty"`
		Name string `yaml:"name,omitempty"`
	} `yaml:"db,omitempty"`
}

func LoadConfig(cPath string) (*Config, error) {
	contents, err := ioutil.ReadFile(cPath)
	if err != nil {
		return nil, err
	}
	c := &Config{}

	if err := yaml.Unmarshal(contents, c); err != nil {
		return nil, fmt.Errorf("load config file err: %s", err)
	}

	return c, nil
}