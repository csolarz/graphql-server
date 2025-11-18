package main

import (
	"log"
	"os"

	"github.com/csolarz/graphql-server/controller"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	engine := controller.StartRouter()

	if err := engine.Run(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
