package config

import (
	"fmt"
	"order-service/pkg/cache"
	"order-service/pkg/db"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server Server            `yaml:"server"`
	DB     db.ConfigDB       `yaml:"db"`
	Redis  cache.ConfigRedis `yaml:"redis"`
}

type Server struct {
	Port        string `yaml:"port" env:"PORT" env-default:"50051"`
	GatewayPort string `yaml:"gateway_port" env:"GATEWAY_PORT" env-default:"8080"`
	Host        string `yaml:"host" env:"HOST" env-default:"localhost"`
	Timeout     int    `yaml:"timeout" env:"TIMEOUT" env-default:"5"`
	BaseUrl     string `yaml:"base_url" env:"BASE_URL" env-default:"http://localhost:50051"`
	Env         string `yaml:"env" env:"ENV" env-default:"production"`
}

func MustParseConfig(path string) (*Config, error) {
	cfg := &Config{}
	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		return nil, fmt.Errorf("error parsing config: %w", err)
	}
	return cfg, nil
}
