package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartRouter() {
	r := gin.Default()

	dependencies := registerDependencies()

	r.GET("/graphql/", dependencies.GraphQLController.Playground())
	r.POST("/query", dependencies.GraphQLController.Query())

	r.POST("/loans", dependencies.ApiController.NewLoan)
	r.GET("/loans/:id", dependencies.ApiController.GetLoan)

	r.POST("/users", dependencies.ApiController.NewUser)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome to the Demo Server")
	})
	r.GET("/ping", pingController)

	r.Run()
}

func pingController(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
