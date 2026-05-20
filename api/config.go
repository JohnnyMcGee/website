package main

import (
	"os"
)

type Config struct {
	Version string
	Host    string
	Port    string
	LogFile string
}

func NewConfig(getEnv func(key string) string) (Config, error) {
	config := Config{
		Version: os.Getenv("VERSION"),
		Host:    os.Getenv("HOST"),
		Port:    os.Getenv("PORT"),
		LogFile: os.Getenv("LOG_FILE"),
	}

	if config.Host == "" {
		config.Host = "localhost"
	}
	if config.Port == "" {
		config.Port = "8000"
	}
	if config.LogFile == "" {
		config.LogFile = "server.log"
	}
	return config, nil
}
