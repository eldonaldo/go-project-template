package core

import (
	"github.com/caarlos0/env/v6"
	"log"
)

const (
	DatabaseDriver = "sqlite3"
	DatabaseFile   = "database.db"
)

// Config elements which can be configured using environment variables
type Config struct {
	Port int `env:"PORT" envDefault:"8008"`
}

// Parse config from environment variables
func ReadConfig() Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Printf("%+v\n", err)
	}

	return cfg
}
