// ...existing code...
package graph

import (
	"context"
	"errors"
	"testing"

	"github.com/csolarz/graphql-server/graph/model"
	"github.com/stretchr/testify/assert"
)

type mockService struct{}

func (m *mockService) Users(ctx context.Context) ([]*model.User, error) {
	return []*model.User{{ID: "1", Name: "Test User"}}, nil
}
func (m *mockService) Payments(ctx context.Context) ([]*model.Payment, error) {
	return []*model.Payment{{ID: "1", Amount: 100}}, nil
}
func (m *mockService) CreatePayment(ctx context.Context, input model.NewPayment) (*model.Payment, error) {
	if input.Amount <= 0 {
		return nil, errors.New("invalid amount")
	}
	return &model.Payment{ID: "2", Amount: input.Amount}, nil
}

func newTestResolver() *Resolver {
	return &Resolver{Service: &mockService{}}
}

func Test_queryResolver_Users(t *testing.T) {
	r := &queryResolver{newTestResolver()}
	users, err := r.Users(context.Background())
	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, "Test User", users[0].Name)
}

func Test_queryResolver_Payments(t *testing.T) {
	r := &queryResolver{newTestResolver()}
	payments, err := r.Payments(context.Background())
	assert.NoError(t, err)
	assert.Len(t, payments, 1)
	assert.Equal(t, float64(100), payments[0].Amount)
}

func Test_mutationResolver_CreatePayment(t *testing.T) {
	r := &mutationResolver{newTestResolver()}
	payment, err := r.CreatePayment(context.Background(), model.NewPayment{Amount: 50})
	assert.NoError(t, err)
	assert.Equal(t, float64(50), payment.Amount)

	_, err = r.CreatePayment(context.Background(), model.NewPayment{Amount: 0})
	assert.Error(t, err)
}
