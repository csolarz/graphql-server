package controller

import (
	"github.com/csolarz/graphql-server/graph"
	"github.com/csolarz/graphql-server/infraestructure/document"
	"github.com/csolarz/graphql-server/usecase/api"
	"github.com/csolarz/graphql-server/usecase/graphql"
)

type dependencies struct {
	GraphQLController *GraphQLController
	ApiController     *ApiController
}

func registerDependencies() dependencies {
	dynamo := document.NewDynamoImpl()

	// GraphQL dependencies
	graphService := graphql.NewService(dynamo)
	graphResolver := graph.NewResolver(graphService)
	graphCtrl := NewGraphQLController(graphResolver)

	// Loan dependencies
	loanService := api.NewService(dynamo)
	loanCtrl := NewApiController(loanService)

	return dependencies{
		GraphQLController: graphCtrl,
		ApiController:     loanCtrl,
	}
}
