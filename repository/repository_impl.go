package repository

import (
	"context"
	"time"

	"github.com/csolarz/graphql-server/graph/model"
	"github.com/csolarz/graphql-server/infraestructure"
)

type RepositoryImpl struct {
	dynamo infraestructure.Dynamo
}

func NewRepositoryImpl(dynamo infraestructure.Dynamo) *RepositoryImpl {
	return &RepositoryImpl{dynamo: dynamo}
}

func (r *RepositoryImpl) Users(ctx context.Context, key string) (*model.User, error) {
	var user *model.User
	err := r.dynamo.Get(ctx, "Users", key, &user)
	return user, err
}

func (r *RepositoryImpl) Payments(ctx context.Context, key string) (*model.Payment, error) {
	var payment *model.Payment
	err := r.dynamo.Get(ctx, "Payments", key, &payment)
	return payment, err
}

func (r *RepositoryImpl) CreatePayment(ctx context.Context, input model.NewPayment) (*model.Payment, error) {
	payment := &model.Payment{
		ID:     time.Now().Unix(), // Ajusta según tu modelo
		Amount: input.Amount,
		User:   &model.User{ID: input.UserID, Name: "Test User"}, // Mock user, ajusta según sea necesario
	}

	err := r.dynamo.Set(ctx, "Payments", payment)
	return payment, err
}
