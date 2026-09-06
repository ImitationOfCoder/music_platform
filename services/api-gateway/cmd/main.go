package main

import (
	"api_gateway_microservice/config"
	"api_gateway_microservice/internal/app"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config error: %s", err)
	}

	app.Run(cfg)
}
