package document

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

type CosmosImpl struct {
	client   *azcosmos.Client
	database string
}

// NewCosmosImpl crea el cliente y lee las variables de entorno
func NewCosmosImpl() *CosmosImpl {
	endpoint := getEnv("COSMOSDB_ENDPOINT", "")
	key := getEnv("COSMOSDB_KEY", "")
	dbName := getEnv("COSMOSDB_DATABASE", "appdb")

	cred, err := azcosmos.NewKeyCredential(key)
	if err != nil {
		panic(err)
	}

	client, err := azcosmos.NewClientWithKey(endpoint, cred, nil)
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

	// Verificamos que el JSON incluya "id"
	var tmp map[string]any
	if err := json.Unmarshal(body, &tmp); err != nil {
		return err
	}

	idVal, ok := tmp["id"].(string)
	if !ok {
		return fmt.Errorf("data struct must contain field `id` as string")
	}

	pk := azcosmos.NewPartitionKeyString(idVal)

	_, err = container.UpsertItem(ctx, pk, body, nil)
	return err
}
