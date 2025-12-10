package document

import "context"

// KeyStore define los métodos para conectar a una base de datos NoSQL tipo key-value.
//
//go:generate mockery --name=KeyStore --output=./mock --outpkg=mock --case=snake
type KeyStore interface {
	Get(ctx context.Context, table string, id string, out any) error
	Set(ctx context.Context, table string, data any) error
}
