package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	MongoHost     string `envconfig:"MONGO_HOST" default:"localhost"`
	MongoDatabase string `envconfig:"MONGO_DB" default:"lifeplan"`
	MongoUser     string `envconfig:"MONGO_USER" default:""`
	MongoPassword string `envconfig:"MONGO_PASSWORD" default:""`
	Environment   string `envconfig:"ENVIRONMENT" default:"dev"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("", cfg)

	return cfg, err
}
