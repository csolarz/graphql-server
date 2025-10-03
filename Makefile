# Makefile
mocks:
	go generate ./...∫

graph-model:
	go run github.com/99designs/gqlgen generate .

################################################
# Run CI-like checks locally

test:
	go test -race -covermode=atomic -v ./... -coverprofile=coverage.out

cover: test
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html

lint:
	golangci-lint run;

sast:
	gosec -exclude-generated -exclude=G115,G401 ./...

################################################
# Running locally with DynamoDB Local

start-local:
	docker compose up -d
	sleep 3
	make create-local-tables

create-local-tables:
	aws dynamodb create-table \
		--table-name Payments \
		--attribute-definitions AttributeName=id,AttributeType=N \
		--key-schema AttributeName=id,KeyType=HASH \
		--provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 \
		--endpoint-url http://localhost:8000 \
		--region us-west-2
	aws dynamodb create-table \
		--table-name Users \
		--attribute-definitions AttributeName=id,AttributeType=N \
		--key-schema AttributeName=id,KeyType=HASH \
		--provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 \
		--endpoint-url http://localhost:8000 \
		--region us-west-2;

# Optional: delete inmemory tables
delete-local-tables:
	aws dynamodb delete-table --table-name Payments --endpoint-url http://localhost:8000 --region us-west-2 || true
	aws dynamodb delete-table --table-name Users --endpoint-url http://localhost:8000 --region us-west-2 || true;

