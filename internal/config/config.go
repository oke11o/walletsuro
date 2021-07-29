package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	PgDSN  string `envconfig:"pg_dsn"`
	DbName string `envconfig:"db_name"`
}

func NewFromEnv() (Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	return cfg, err
}
