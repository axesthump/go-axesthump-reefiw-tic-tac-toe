package main

import (
	"fmt"
	"go-axesthump-reefiw-tic-tac-toe/internal/config"
	"log"
	"net/http"
	"os"
)

func main() {
	conf, err := config.NewAppConfig()
	if err != nil {
		exit(err.Error())
	}
	log.Printf("Start listen server at %s\n", conf.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), conf.Handler.Router)
	if err != nil {
		exit(err.Error())
	}
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
