package service

import (
	"context"
	"testing"

	"github.com/csolarz/graphql-server/graph/model"
	"github.com/stretchr/testify/assert"
)

func TestService_Users(t *testing.T) {
	s := NewService()
	_, err := s.Users(context.Background())
	assert.Error(t, err)
}

func TestService_Payments(t *testing.T) {
	s := NewService()
	_, err := s.Payments(context.Background())
	assert.Error(t, err)
}

func TestService_CreatePayment(t *testing.T) {
	s := NewService()
	input := model.NewPayment{}
	_, err := s.CreatePayment(context.Background(), input)
	assert.Error(t, err)
}
