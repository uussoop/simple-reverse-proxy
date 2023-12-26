package configreader

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Proxies []Proxy `yaml:"proxies"`
}

type Proxy struct {
	Name string `yaml:"name"`
	From Url    `yaml:"from"`
	To   Url    `yaml:"to"`
}
type Url struct {
	Scheme string `yaml:"scheme"`
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
}

func ReadConfig(filename string) (*Config, error) {
	var config Config

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
