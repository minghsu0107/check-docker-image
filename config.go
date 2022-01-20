package main

import "github.com/kelseyhightower/envconfig"

// Config defines all needed parameters
type Config struct {
	Registry *RegistryConfig
	Images   []string `envconfig:"CHECKED_IMAGES"`
	LogLevel string   `envconfig:"LOGLEVEL"`
}

// RegistryConfig details
type RegistryConfig struct {
	Url      string `envconfig:"REGISTRY_URL"`
	Username string `envconfig:"REGISTRY_USERNAME"`
	Password string `envconfig:"REGISTRY_PASSWORD"`
}

func readEnv(config *Config) error {
	err := envconfig.Process("plugin", config)
	if err != nil {
		return err
	}
	return nil
}

func setDefault(config *Config) {
	config.Registry = &RegistryConfig{}
	// log level can be debug, info, warn, error
	config.LogLevel = "info"
}
