package apiserver

import (
	"github.com/MirrexOne/http-rest-api/internal/app/store/sqlstore"
)

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *sqlstore.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    sqlstore.NewConfig(),
	}
}
