package main

import (
	"log"

	"github.com/enyasantos/go-async-order-system/config"
	"github.com/enyasantos/go-async-order-system/consumer"
	"github.com/enyasantos/go-async-order-system/router"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Inicializa configurações do aplicativo
	if err := config.Init(); err != nil {
		log.Fatalf("Config initialization error: %v", err)
	}

	// Inicializa o roteamento HTTP
	go func() {
		router.Initialize()
	}()

	consumer.InitializeConnectionChannelRabbitMQ()

}
