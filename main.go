package main

import (
	"log"
	"os"

	"github.com/csolarz/graphql-server/controller"
	"github.com/csolarz/graphql-server/infraestructure/logger"
)

const defaultPort = "8080"

func main() {
	logger.Init()
	defer logger.Sync()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	engine := controller.StartRouter()

	if err := engine.Run(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
