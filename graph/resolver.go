package graph

import "github.com/csolarz/graphql-server/usecase/graphql"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Service graphql.ResolverUsecase
}

func NewResolver(service graphql.ResolverUsecase) Resolver {
	return Resolver{Service: service}
}
