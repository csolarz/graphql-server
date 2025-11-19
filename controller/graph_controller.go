package controller

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/csolarz/graphql-server/graph"
	"github.com/csolarz/graphql-server/graph/generated"
	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/v2/ast"
)

type GraphQLController struct {
	service graph.Resolver
}

func NewGraphQLController(service graph.Resolver) *GraphQLController {
	return &GraphQLController{service: service}
}

func (gc GraphQLController) Query() gin.HandlerFunc {

	h := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: &gc.service}))

	// Server setup:
	h.AddTransport(transport.Options{})
	h.AddTransport(transport.GET{})
	h.AddTransport(transport.POST{})

	h.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	h.Use(extension.Introspection{})
	h.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	return func(ctx *gin.Context) {
		h.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func (gc GraphQLController) Playground() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
