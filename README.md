# Configuración de credenciales para DynamoDB Local

Para evitar el error `Unable to locate credentials` al usar AWS CLI o el SDK de AWS con DynamoDB Local, debes configurar credenciales dummy (no reales). Esto es necesario aunque DynamoDB Local no las valide.

Ejecuta en tu terminal:

```sh
aws configure
```

Cuando te pida los datos, puedes poner cualquier valor:

- AWS Access Key ID: `dummy`
- AWS Secret Access Key: `dummy`
- Default region name: `us-west-2`
- Default output format: `json`

Esto crea el archivo `~/.aws/credentials` y permite que el cliente funcione sin errores.
## ¿Cómo generar mocks (sin go:generate)?

Ejecuta el comando mockery directamente en cada package donde tengas interfaces:

```sh
# Ejemplo para la interface Repository
mockery --name=Repository --output=repository/mock --outpkg=mock --case=snake

# Ejemplo para la interface Resolver
mockery --name=Resolver --output=service/mock --outpkg=mock --case=snake

# Ejemplo para Dynamo y DynamoDBAPI
mockery --name=Dynamo --output=infraestructure/mock --outpkg=mock --case=snake
mockery --name=DynamoDBAPI --output=infraestructure/mock --outpkg=mock --case=snake
```

Esto generará los mocks en la carpeta mock de cada package, usando formato snake_case.
## Ejemplos de Queries y Mutations

### Query: Obtener usuario

```graphql
query {
	users {
		id
		name
	}
}
```

### Query: Obtener pago

```graphql
query {
	payments {
		id
		amount
		user {
			id
			name
		}
	}
}
```

### Mutation: Crear pago

```graphql
mutation {
	create_payment(input: { amount: 100, user_id: "1" }) {
		id
		amount
		user {
			id
			name
		}
	}
}
```
# graphql-server
Servidor GraphQL de ejemplo en Go

## Requisitos

- Go 1.25+
- Docker (opcional, para DynamoDB local)

## ¿Cómo ejecutar el servidor?

```sh
go run main.go
```

O bien, usando el Makefile:

```sh
make run
```

## ¿Cómo debuggear?

Puedes usar Delve (dlv) para debuggear:

```sh
dlv debug main.go
```
O desde VS Code usando el launch.json incluido.


## ¿Cómo construir el modelo GraphQL?

Edita el archivo `graph/schema.graphqls` y luego ejecuta:

```sh
make graph-model
```

## ¿Cómo generar mocks?

```sh
make mocks
```
Esto generará los mocks en la carpeta `mock` de cada package con interfaces.

## Como simular CI localmente

### Linter
Lint:
```sh
make lint
```
### SAST (análisis de seguridad):
```sh
make sast
```

### Tests y cobertura

```sh
make test
```
Para ver el reporte de cobertura:

```sh
make cover
```

## ¿Cómo simular el DynamoDB localmente?

Puedes levantar un contenedor local de DynamoDB para desarrollo y testing. Consulta el archivo docker-compose.yml si está presente.

```sh
make start-local // Levanta DynamoDB local y crea las tablas en memoria, sin persitencia
```

