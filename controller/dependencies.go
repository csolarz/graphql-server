package controller

import (
	"os"

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
	keyStore := resolveKeyStore()

	// GraphQL dependencies
	graphService := graphql.NewService(keyStore)
	graphResolver := graph.NewResolver(graphService)
	graphCtrl := NewGraphQLController(graphResolver)

	// Loan dependencies
	loanService := api.NewService(keyStore)
	loanCtrl := NewApiController(loanService)

	return dependencies{
		GraphQLController: graphCtrl,
		ApiController:     loanCtrl,
	}
}

func resolveKeyStore() document.KeyStore {
	if cloudProvider := os.Getenv("CLOUD_PROVIDER"); cloudProvider == "AWS" || cloudProvider == "" {
		return document.NewDynamoImpl()
	}

	return document.NewCosmosImpl()
}
