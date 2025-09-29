package controller

import (
	"github.com/csolarz/graphql-server/graph"
	"github.com/csolarz/graphql-server/service"
)

func registerDependencies() graph.Resolver {
	svc := service.NewService()
	return graph.Resolver{Service: svc}
}
