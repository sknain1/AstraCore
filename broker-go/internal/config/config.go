package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
		Env     string `yaml:"env"`
	} `yaml:"app"`

	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`

	Broker struct {
		Name string `yaml:"name"`
	} `yaml:"broker"`

	Logging struct {
		Level string `yaml:"level"`
	} `yaml:"logging"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
