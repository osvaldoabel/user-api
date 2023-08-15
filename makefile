.PHONY: default run build test docs clean
# Variables
APP_NAME=user-api

# Tasks
default: run-with-docs

run:
	@go run cmd/main.go
run-with-docs:
	@swag init -g cmd/main.go
	@go run cmd/main.go
build:
	@go build -o $(APP_NAME) cmd/main.go
test:
	@go test ./ ...
docs:
	sh ./scripts/generate_swagger_docs.sh
