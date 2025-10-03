package infraestructure

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/stretchr/testify/assert"
)

type mockDynamoDBClient struct {
	GetItemFunc func(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
	PutItemFunc func(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
}

func (m *mockDynamoDBClient) GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	return m.GetItemFunc(ctx, params, optFns...)
}

func (m *mockDynamoDBClient) PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	return m.PutItemFunc(ctx, params, optFns...)
}

func TestDynamoImpl_Get_Success(t *testing.T) {
	mockClient := &mockDynamoDBClient{
		GetItemFunc: func(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
			return &dynamodb.GetItemOutput{
				Item: map[string]types.AttributeValue{
					"id": &types.AttributeValueMemberS{Value: "1"},
				},
			}, nil
		},
	}
	d := NewDynamoImplWithClient(mockClient)
	var out map[string]interface{}
	err := d.Get(context.Background(), "Users", "1", &out)
	assert.NoError(t, err)
	assert.Equal(t, "1", out["id"])
}

func TestDynamoImpl_Get_NotFound(t *testing.T) {
	mockClient := &mockDynamoDBClient{
		GetItemFunc: func(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
			return &dynamodb.GetItemOutput{Item: nil}, nil
		},
	}
	d := NewDynamoImplWithClient(mockClient)
	var out map[string]interface{}
	err := d.Get(context.Background(), "Users", "1", &out)
	assert.Error(t, err)
}

func TestDynamoImpl_Set_Success(t *testing.T) {
	mockClient := &mockDynamoDBClient{
		PutItemFunc: func(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
			return &dynamodb.PutItemOutput{}, nil
		},
	}
	d := NewDynamoImplWithClient(mockClient)
	data := map[string]interface{}{"id": "1"}
	err := d.Set(context.Background(), "Users", data)
	assert.NoError(t, err)
}

func TestDynamoImpl_Set_Error(t *testing.T) {
	mockClient := &mockDynamoDBClient{
		PutItemFunc: func(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
			return nil, errors.New("put error")
		},
	}
	d := NewDynamoImplWithClient(mockClient)
	data := map[string]interface{}{"id": "1"}
	err := d.Set(context.Background(), "Users", data)
	assert.Error(t, err)
}
