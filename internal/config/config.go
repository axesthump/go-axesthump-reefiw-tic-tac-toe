package config

import (
	"flag"
	"go-axesthump-reefiw-tic-tac-toe/internal/handlers"
)

type AppConfig struct {
	Port    string
	Handler *handlers.AppHandler
}

func NewAppConfig() (*AppConfig, error) {
	port := flag.String("p", "8080", "port")
	flag.Parse()
	handler, err := handlers.NewAppHandler()
	if err != nil {
		return nil, err
	}
	conf := &AppConfig{
		Port:    *port,
		Handler: handler,
	}
	return conf, err
}
