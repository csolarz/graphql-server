package api

import (
	"context"
	"fmt"
	"time"

	"github.com/csolarz/graphql-server/entities"
	"github.com/csolarz/graphql-server/infraestructure/document"
)

const LoansTable = "Loans"
const UsersTable = "Users"

type Service struct {
	repo document.KeyStore
}

//go:generate mockery --name=Usecase --output=./mock --outpkg=mock --case=snake
type Usecase interface {
	NewLoan(ctx context.Context, request entities.LoanRequest) (*entities.Loan, error)
	GetLoan(ctx context.Context, loanID int64) (*entities.Loan, error)
	NewUser(ctx context.Context, request entities.UserRequest) (*entities.User, error)
}

func NewService(repo document.KeyStore) *Service {
	return &Service{repo: repo}
}

func (s *Service) NewLoan(ctx context.Context, request entities.LoanRequest) (*entities.Loan, error) {
	loan := entities.NewLoan(request)

	err := s.repo.Set(ctx, LoansTable, loan)
	if err != nil {
		return nil, err
	}

	return loan, nil
}

func (s *Service) GetLoan(ctx context.Context, loanID int64) (*entities.Loan, error) {
	var loan *entities.Loan

	err := s.repo.Get(ctx, LoansTable, fmt.Sprint(loanID), &loan)
	if err != nil {
		return nil, err
	}

	return loan, nil
}

func (s *Service) NewUser(ctx context.Context, request entities.UserRequest) (*entities.User, error) {
	user := entities.User{
		ID:    int64(time.Now().Second()),
		Name:  request.Name,
		Loans: []entities.Loan{},
	}

	err := s.repo.Set(ctx, UsersTable, user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
