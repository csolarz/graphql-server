package main

import (
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

	//nolint
	engine.Run(":" + port)
}
