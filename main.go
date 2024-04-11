package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pe-Gomes/short-url/api"
	db "github.com/pe-Gomes/short-url/infra/db/repository"
	"github.com/pe-Gomes/short-url/util"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("Error loading enviroment variables: ", err)
	}

	connPoll, err := pgxpool.New(context.Background(), config.DatabaseURL)

	if err != nil {
		log.Fatal("Could not create connection pool: ", err)
	}

	store := db.NewStore(connPoll)

	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal("Could not create server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Could not start server: ", err)
	}
}
