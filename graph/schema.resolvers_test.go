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

func (m *mockService) Users(ctx context.Context, id int64) (*model.User, error) {
	return &model.User{ID: id, Name: "Test User"}, nil
}
func (m *mockService) Payments(ctx context.Context, id int64) (*model.Payment, error) {
	return &model.Payment{ID: id, Amount: 100, User: &model.User{ID: id, Name: "Test User"}}, nil
}
func (m *mockService) CreatePayment(ctx context.Context, input model.NewPayment) (*model.Payment, error) {
	if input.Amount <= 0 {
		return nil, errors.New("invalid amount")
	}
	// Convierte el UserID string a int64 para el mock
	var uid int64 = 1
	return &model.Payment{ID: 2, Amount: input.Amount, User: &model.User{ID: uid, Name: "Test User"}}, nil
}

func newTestResolver() *Resolver {
	return &Resolver{Service: &mockService{}}
}

func Test_queryResolver_User(t *testing.T) {
	r := &queryResolver{newTestResolver()}
	user, err := r.User(context.Background(), int64(1))
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "Test User", user.Name)
}

func Test_queryResolver_Payment(t *testing.T) {
	r := &queryResolver{newTestResolver()}
	payment, err := r.Payment(context.Background(), int64(1))
	assert.NoError(t, err)
	assert.NotNil(t, payment)
	assert.Equal(t, float64(100), payment.Amount)
	assert.NotNil(t, payment.User)
	assert.Equal(t, "Test User", payment.User.Name)
}

func Test_mutationResolver_CreatePayment(t *testing.T) {
	r := &mutationResolver{newTestResolver()}
	payment, err := r.CreatePayment(context.Background(), model.NewPayment{Amount: 50})
	assert.NoError(t, err)
	assert.Equal(t, float64(50), payment.Amount)

	_, err = r.CreatePayment(context.Background(), model.NewPayment{Amount: 0})
	assert.Error(t, err)
}
