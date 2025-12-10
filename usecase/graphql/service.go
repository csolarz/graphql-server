package graphql

import (
	"context"
	"fmt"

	"github.com/csolarz/graphql-server/graph/model"
	"github.com/csolarz/graphql-server/infraestructure/document"
)

// GraphQLService define la interfaz principal para los servicios de GraphQL
//
//go:generate mockery --name=ResolverUsecase --output=./mock --outpkg=mock --case=snake
type ResolverUsecase interface {
	// Users obtiene un usuario por ID
	User(ctx context.Context, id int64) (*model.User, error)

	// Loans obtiene un préstamo por ID
	Loan(ctx context.Context, id int64) (*model.Loan, error)

	// Installment obtiene una cuota por ID
	Installment(ctx context.Context, id int64) (*model.Installment, error)
}

type Service struct {
	repo document.KeyStore
}

func NewService(repo document.KeyStore) *Service {
	return &Service{repo: repo}
}

func (s *Service) User(ctx context.Context, id int64) (*model.User, error) {
	var user *model.User
	err := s.repo.Get(ctx, "Users", fmt.Sprintf("%d", id), &user)

	return user, err
}

func (s *Service) Loan(ctx context.Context, id int64) (*model.Loan, error) {
	var loan model.Loan
	err := s.repo.Get(ctx, "Loans", fmt.Sprintf("%d", id), &loan)

	return &loan, err
}

func (s *Service) Installment(ctx context.Context, id int64) (*model.Installment, error) {
	var installment *model.Installment
	err := s.repo.Get(ctx, "Installments", fmt.Sprintf("%d", id), &installment)

	return installment, err
}
