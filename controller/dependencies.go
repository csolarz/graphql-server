package controller

import (
	"github.com/csolarz/graphql-server/graph"
	"github.com/csolarz/graphql-server/infraestructure"
	"github.com/csolarz/graphql-server/repository"
	"github.com/csolarz/graphql-server/service"
)

func registerDependencies() graph.Resolver {
	dynamo := infraestructure.NewDynamoImpl()
	repo := repository.NewRepositoryImpl(dynamo)
	svc := service.NewService(repo)
	return graph.Resolver{Service: svc}
}
