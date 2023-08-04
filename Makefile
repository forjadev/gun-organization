.PHONY: default run build test docs clean

#Variables
APP_NAME=gunorganization

#Tasks
default: run-with-docs

run: database
	@go run main.go
run-with-docs: database
	@swag init
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./...
docs:
	@swag init
database:
	@docker compose down
	@docker compose up -d
	@sleep 0.5
clean:
	@rm -rf $(APP_NAME)
	@rm -rf ./docs
	@docker compose down -v