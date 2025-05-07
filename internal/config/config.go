package config

import (
	"os"
	"sync"
)

type Config struct {
	App      Application
	Postgres Database
	Mongo    Service
	Redis    Service
	Kafka    Service
}

type Application struct {
	Name        string
	Environment string
	Port        string
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type Service struct {
	Host string
	Port string
}

var (
	instance *Config
	once     sync.Once
)

func Load() *Config {
	once.Do(func() {
		instance = &Config{
			App: Application{
				Name:        os.Getenv("APP_NAME"),
				Environment: os.Getenv("APP_ENV"),
				Port:        os.Getenv("APP_PORT"),
			},
			Postgres: Database{
				Host:     os.Getenv("DB_HOST"),
				Port:     os.Getenv("DB_PORT"),
				User:     os.Getenv("DB_USER"),
				Password: os.Getenv("DB_PASSWORD"),
				Name:     os.Getenv("DB_NAME"),
			},
			Redis: Service{
				Host: os.Getenv("REDIS_HOST"),
				Port: os.Getenv("REDIS_PORT"),
			},
			Mongo: Service{
				Host: os.Getenv("MONGO_HOST"),
				Port: os.Getenv("MONGO_PORT"),
			},
			Kafka: Service{
				Host: os.Getenv("KAFKA_HOST"),
				Port: os.Getenv("KAFKA_PORT"),
			},
		}
	})

	return instance
}
