//go:generate mockery --name=Repository --output=./mock --outpkg=mock --case=snake
package repository

import (
	"context"

	"github.com/csolarz/graphql-server/graph/model"
)

type Repository interface {
	Users(ctx context.Context, key string) (*model.User, error)
	Payments(ctx context.Context, key string) (*model.Payment, error)
	CreatePayment(ctx context.Context, input model.NewPayment) (*model.Payment, error)
}
