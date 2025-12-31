package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	BaseUrl string `yaml:"base_url" env:"BASE_URL" env-default:"http://localhost:8080"`
	Server  Server
}

type Server struct {
	Port    string `yaml:"port" env:"PORT" env-default:"8080"`
	Host    string `yaml:"host" env:"HOST" env-default:"localhost"`
	Timeout int    `yaml:"timeout" env:"TIMEOUT" env-default:"5"`
}

func MustParseConfig(path string) (*Config, error) {
	cfg := &Config{}
	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		return nil, fmt.Errorf("error parsing config: %w", err)
	}
	return cfg, nil
}
