//go:generate mockery --name=Resolver --output=./mock --outpkg=mock --case=snake
package service

import (
	"context"

	"github.com/csolarz/graphql-server/graph/model"
	"github.com/csolarz/graphql-server/repository"
)

// GraphQLService define la interfaz principal para los servicios de GraphQL
type Resolver interface {
	Users(ctx context.Context) (*model.User, error)
	Payments(ctx context.Context) (*model.Payment, error)
	CreatePayment(ctx context.Context, input model.NewPayment) (*model.Payment, error)
}

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Users(ctx context.Context) (*model.User, error) {
	// Si tu lógica requiere un ID fijo, cámbialo aquí o ajusta el repo
	return s.repo.Users(ctx, "1")
}

func (s *Service) Payments(ctx context.Context) (*model.Payment, error) {
	// Si tu lógica requiere un ID fijo, cámbialo aquí o ajusta el repo
	return s.repo.Payments(ctx, "1")
}

func (s *Service) CreatePayment(ctx context.Context, input model.NewPayment) (*model.Payment, error) {
	return s.repo.CreatePayment(ctx, input)
}
