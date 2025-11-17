# Makefile
mocks:
	go generate ./...

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
	docker-compose up -d
	scripts/dynamo-local.sh


# Optional: delete inmemory tables
delete-tables:
	aws dynamodb delete-table --table-name Installments --endpoint-url http://localhost:8000 --region us-west-2 || true;
	aws dynamodb delete-table --table-name Users --endpoint-url http://localhost:8000 --region us-west-2 || true;
	aws dynamodb delete-table --table-name Loans --endpoint-url http://localhost:8000 --region us-west-2 || true;

