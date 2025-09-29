package service

import (
	"context"
	"fmt"

	"github.com/csolarz/graphql-server/graph/model"
)

// GraphQLService define la interfaz principal para los servicios de GraphQL
type Resolver interface {
	Users(ctx context.Context) ([]*model.User, error)
	Payments(ctx context.Context) ([]*model.Payment, error)
	CreatePayment(ctx context.Context, input model.NewPayment) (*model.Payment, error)
}

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Users(ctx context.Context) ([]*model.User, error) {
	return nil, fmt.Errorf("not implemented: Users - users")
}

func (s *Service) Payments(ctx context.Context) ([]*model.Payment, error) {
	return nil, fmt.Errorf("not implemented: Payments - payments")
}

func (s *Service) CreatePayment(ctx context.Context, input model.NewPayment) (*model.Payment, error) {
	return nil, fmt.Errorf("not implemented: CreatePayment - createPayment")
}
