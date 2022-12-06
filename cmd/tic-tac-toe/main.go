package main

import (
	"fmt"
	"go-axesthump-reefiw-tic-tac-toe/internal/config"
	"log"
	"net/http"
	"os"
)

func main() {
	conf := config.NewAppConfig()
	log.Printf("Start listen server at %s\n", conf.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), conf.Handler.Router)
	if err != nil {
		log.Printf("Cant start server... %s\n", err.Error())
		os.Exit(1)
	}
}
