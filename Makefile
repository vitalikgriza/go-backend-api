build:
	@go build -o bin/go-backend-api

run: build
	@./bin/go-backend-api

test:
	@go test -v ./...