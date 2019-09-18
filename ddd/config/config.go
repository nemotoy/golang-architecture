package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Config struct {
	Port int `required:"true"`
}

func NewConfig() (*Config, error) {
	var c Config
	if err := envconfig.Process("myapp", &c); err != nil {
		return nil, errors.WithStack(err)
	}
	return &c, nil
}
