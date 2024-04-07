package main

import (
	"nestle/internal/api"
	"nestle/internal/api/router"
	"nestle/internal/config"
	"nestle/internal/database"
	"nestle/internal/logs"
)

func main() {
	logger := logs.InitLogger()
	config, err := config.Load("./")
	if err != nil {
		logger.Fatal(err.Error())
	}
	repo := database.InitRepository(&config.Database, logger)
	server := api.NewServer(config.Server)
	server.RegisterDb(repo)
	router.Init(server)
	server.StartServer()
}
