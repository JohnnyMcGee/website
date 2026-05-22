package config

import (
	"errors"
)

type Config struct {
	Version         string
	Host            string
	Port            string
	LogFile         string
	ResendApiKey    string
	ContactEmail    string
	ResendFromEmail string
}

var (
	ErrMissingResendApiKey = errors.New("missing RESEND_API_KEY environment variable")
	ErrMissingContactEmail = errors.New("missing CONTACT_EMAIL environment variable")
)

func New(getEnv func(key string) string) (Config, error) {
	config := Config{
		ContactEmail:    getEnv("CONTACT_EMAIL"),
		Version:         getEnv("VERSION"),
		Host:            getEnv("HOST"),
		Port:            getEnv("PORT"),
		LogFile:         getEnv("LOG_FILE"),
		ResendApiKey:    getEnv("RESEND_API_KEY"),
		ResendFromEmail: getEnv("RESEND_FROM_EMAIL"),
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
	if config.ResendApiKey == "" {
		return Config{}, ErrMissingResendApiKey
	}
	if config.ContactEmail == "" {
		return Config{}, ErrMissingContactEmail
	}
	if config.ResendFromEmail == "" {
		config.ResendFromEmail = "onboarding@resend.dev"
	}
	return config, nil
}

type ResendNotifierConfig struct {
	ToEmail   string
	FromEmail string
}

func (c *Config) ResendNotifier() ResendNotifierConfig {
	return ResendNotifierConfig{
		ToEmail:   c.ContactEmail,
		FromEmail: c.ResendFromEmail,
	}
}
