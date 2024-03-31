package config

import (
	"github.com/caarlos0/env/v10"
	"go.uber.org/fx"
)

// Module provides a Config struct with values from environment variables.
var Module = fx.Provide(New)

// Config holds configuration for the application.
type Config struct {
	AppMode string         `env:"APP_MODE"              envDefault:"development"`
	Port    string         `env:"PORT"                  envDefault:"8000"`
	DBCfg   DataBaseConfig `envPrefix:"DB_"`
}

type DataBaseConfig struct {
	Port     string `env:"PORT"         envDefault:"5432"`
	Host     string `env:"HOST"         envDefault:"localhost"`
	User     string `env:"USER"         envDefault:"postgres"`
	Password string `env:"PASSWORD"     envDefault:"postgres"`
	Database string `env:"DATABASE"     envDefault:"advertise"`
}

// New returns a new Config struct with values from environment variables.
func New() *Config {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}
	return &cfg
}
