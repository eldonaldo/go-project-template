package core

import (
	"github.com/caarlos0/env/v6"
	"log"
)

// Not configurable - internal only - config
const (
	DatabaseDriver = "sqlite3"
	DatabaseFile   = "database.db"
)

// Config for which the fields can be configured using environment variables
type Config struct {

	// Reads the content of environment variable PORT as int into this field.
	// If none is setup, the default value 8008 is used
	Port int `env:"PORT" envDefault:"8008"`
}

// Parse config from environment variables, if possible
func ReadConfig() Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Printf("%+v\n", err)
	}

	return cfg
}
