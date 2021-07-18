package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	PgDSN string `envconfig:"pg_dsn"`
}

func NewFromEnv() (Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	return cfg, err
}
