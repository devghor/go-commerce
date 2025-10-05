package main

import (
	"go-commerce/configs"
	"go-commerce/internal/api"
	"log"
)

func main() {

	cfg, err := configs.SetupEnv()
	if err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	api.StartServer(cfg)

}
