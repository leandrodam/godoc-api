package config

import "os"

type Config struct {
	Port string
}

func Load() *Config {
	return &Config{
		Port: os.Getenv("APP_PORT"),
	}
}
