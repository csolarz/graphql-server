//go:generate mockery --name=Dynamo --output=./mock --outpkg=mock --case=snake
//go:generate mockery --name=DynamoDBAPI --output=./mock --outpkg=mock --case=snake
package infraestructure

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Dynamo define los métodos expuestos por DynamoImpl (alias interno, si se requiere)
type Dynamo interface {
	Get(ctx context.Context, table string, id string, out any) error
	Set(ctx context.Context, table string, data any) error
}

type DynamoDBAPI interface {
	GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
}

type DynamoImpl struct {
	db DynamoDBAPI
}

func NewDynamoImpl() *DynamoImpl {
	region := getEnv("DYNAMO_REGION", "us-west-2")
	endpoint := getEnv("DYNAMO_ENDPOINT", "http://localhost:8000")
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		//nolint
		config.WithEndpointResolver(
			aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{URL: endpoint, SigningRegion: region, HostnameImmutable: true}, nil
			}),
		),
	)
	if err != nil {
		panic(fmt.Sprintf("unable to load SDK config, %v", err))
	}

	db := dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.Credentials = credentials.NewStaticCredentialsProvider("dummy", "dummy", "")
	})

	return &DynamoImpl{
		db: db,
	}
}

// NewDynamoImplWithClient permite inyectar un cliente custom (para tests)
func NewDynamoImplWithClient(client DynamoDBAPI) *DynamoImpl {
	return &DynamoImpl{db: client}
}

// Get obtiene un item de DynamoDB y lo deserializa en out.
func (r *DynamoImpl) Get(ctx context.Context, table string, id string, out any) error {
	key := map[string]types.AttributeValue{
		"_id": &types.AttributeValueMemberS{Value: fmt.Sprintf("%s_%s", table, id)},
	}

	input := &dynamodb.GetItemInput{
		TableName: &table,
		Key:       key,
	}

	result, err := r.db.GetItem(ctx, input)
	if err != nil {
		return err
	}

	if result.Item == nil {
		return nil
	}

	// Deserializar el resultado
	err = attributevalue.UnmarshalMap(result.Item, out)
	if err != nil {
		return err
	}

	return nil
}

// Set almacena un item en DynamoDB usando el partition key explícito en la data struct.
func (r *DynamoImpl) Set(ctx context.Context, table string, data any) error {
	av, err := attributevalue.MarshalMap(data)
	av["_id"] = &types.AttributeValueMemberS{Value: fmt.Sprintf("%s_%s", table, av["id"].(*types.AttributeValueMemberN).Value)}

	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		TableName: &table,
		Item:      av,
	}
	_, err = r.db.PutItem(ctx, input)
	return err
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
