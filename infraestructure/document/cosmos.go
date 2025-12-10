package document

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

type CosmosImpl struct {
	client   *azcosmos.Client
	database string
}

// NewCosmosImpl crea el cliente y lee las variables de entorno
func NewCosmosImpl() *CosmosImpl {
	endpoint := getEnv("COSMOSDB_ENDPOINT", "")
	dbName := getEnv("COSMOSDB_DATABASE", "appdb")

	// Autenticación con Managed Identity o Azure AD
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		panic(err)
	}

	client, err := azcosmos.NewClient(endpoint, cred, nil)
	if err != nil {
		panic(err)
	}

	return &CosmosImpl{
		client:   client,
		database: dbName,
	}
}

// Para test: cliente inyectable
func NewCosmosImplWithClient(client *azcosmos.Client, database string) *CosmosImpl {
	return &CosmosImpl{
		client:   client,
		database: database,
	}
}

// ---------------------------
// Implementación de la interfaz KeyValue
// ---------------------------

// Get obtiene un documento desde CosmosDB
func (c *CosmosImpl) Get(ctx context.Context, table string, id string, out any) error {
	db, err := c.client.NewDatabase(c.database)
	if err != nil {
		return err
	}

	container, err := db.NewContainer(table)
	if err != nil {
		return err
	}

	pk := azcosmos.NewPartitionKeyString(id)

	resp, err := container.ReadItem(ctx, pk, id, nil)
	if err != nil {
		// Cosmos usa 404 para “no existe”
		if resp.RawResponse != nil && resp.RawResponse.StatusCode == 404 {
			return nil
		}
		return err
	}

	if resp.Value == nil {
		return nil
	}

	return json.Unmarshal(resp.Value, out)
}

// Set guarda o actualiza un documento en CosmosDB
func (c *CosmosImpl) Set(ctx context.Context, table string, data any) error {
	db, err := c.client.NewDatabase(c.database)
	if err != nil {
		return err
	}

	container, err := db.NewContainer(table)
	if err != nil {
		return err
	}

	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	id, exists := getID(data)
	if !exists {
		return fmt.Errorf("data struct must have an 'ID' field")
	}

	pk := azcosmos.NewPartitionKeyString(fmt.Sprintf("%s_%d", table, id.(int64)))

	_, err = container.UpsertItem(ctx, pk, body, nil)
	return err
}

func getID(obj interface{}) (any, bool) {
	v := reflect.ValueOf(obj)

	// Si es un puntero, obtener su valor
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Debe ser struct
	if v.Kind() != reflect.Struct {
		return nil, false
	}

	// Obtener campo "ID"
	field := v.FieldByName("ID")
	if !field.IsValid() {
		return nil, false
	}

	// Convertir a interface y retornar
	return field.Interface(), true
}
