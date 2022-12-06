package config

import (
	"flag"
	"go-axesthump-reefiw-tic-tac-toe/internal/handlers"
)

type AppConfig struct {
	Port    string
	Handler *handlers.AppHandler
}

func NewAppConfig() *AppConfig {
	port := flag.String("p", "8080", "port")
	flag.Parse()
	conf := &AppConfig{
		Port:    *port,
		Handler: handlers.NewAppHandler(),
	}
	return conf
}
