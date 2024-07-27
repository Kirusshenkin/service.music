package main

import (
	"log"
	"net/http"
	"service.music/internal/app"
	"service.music/internal/config"
	"service.music/internal/database"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	db, err := database.InitPostgresDB(cfg)
	if err != nil {
		log.Fatalf("could not connect to postgres: %v", err)
	}
	defer db.Close()

	redisClient := database.InitRedisClient(cfg)
	defer redisClient.Close()

	router := app.SetupRouter()

	log.Println("Starting server on :", cfg.Server.Port)
	if err := http.ListenAndServe(":"+cfg.Server.Port, router); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
