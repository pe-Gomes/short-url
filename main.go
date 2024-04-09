package main

import (
	"log"

	"github.com/pe-Gomes/short-url/api"
	"github.com/pe-Gomes/short-url/util"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("Error loading enviroment variables: ", err)
	}

	server, err := api.NewServer(config)

	if err != nil {
		log.Fatal("Could not create server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Could not start server: ", err)
	}
}
