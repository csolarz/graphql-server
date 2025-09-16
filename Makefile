graph-model:
	go run github.com/99designs/gqlgen generate .

test:
	go test -race -v ./... -coverprofile=coverage.out

cover-html: test
	go tool cover -html=coverage.out -o coverage.html
	sensible-browser coverage.html