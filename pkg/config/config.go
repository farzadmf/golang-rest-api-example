// Package config contains the configuration.
package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

// Config holds configuration value.
type Config struct {
	DBHost     string `env:"DB_HOST,required"`
	DBUser     string `env:"DB_USER,required"`
	DBPassword string `env:"DB_PASSWORD,required"`
	DBName     string `env:"DB_NAME,required"`
}

// New creates a new config.
func New() (*Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		return nil, fmt.Errorf("error creating config: %w", err)
	}

	return &cfg, nil
}
