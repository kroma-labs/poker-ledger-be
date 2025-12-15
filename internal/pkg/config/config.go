package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/rotisserie/eris"
)

type Config struct {
	Env      string        `envconfig:"ENV" default:"development"`
	Port     string        `envconfig:"PORT" default:"8080"`
	Timeout  time.Duration `envconfig:"TIMEOUT" default:"10s"`
	DBString string        `envconfig:"DB_STRING" required:"true"`
}

func Load() (*Config, error) {
	var c Config
	if err := envconfig.Process("", &c); err != nil {
		return nil, eris.Wrap(err, "error loading ENV")
	}
	return &c, nil
}
