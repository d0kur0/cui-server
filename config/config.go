package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Port int `env:"PORT" envDefault:"5050"`
}

func New() *Config {
	return &Config{}
}

func (config *Config) Init() error {
	return env.Parse(config)
}
