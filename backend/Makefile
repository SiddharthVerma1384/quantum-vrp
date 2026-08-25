-include .env
export

all: build

setup:
	@echo "Configuring git hooks..."
	@git config core.hooksPath .githooks
	@echo "Done. Pre-commit checks will now run before every commit."

build:
	@echo "Building binary..."
	@go build -o bin/api cmd/api/main.go

run:
	@echo "Running local server..."
	@go run cmd/api/main.go

dev:
	@echo "Running hot-reload dev server with air..."
	@air -c .air.toml

test:
	@echo "Running tests..."
	@go test -v ./...

fmt:
	@echo "Formatting code..."
	@go fmt ./...

lint:
	@echo "Running golangci-lint..."
	@golangci-lint run --config=.golangci.yml

vulncheck:
	@echo "Running govulncheck..."
	@govulncheck ./...

docker-up:
	@echo "Starting docker services..."
	@docker compose up --build -d

docker-down:
	@echo "Stopping docker services..."
	@docker compose down -v

sqlc:
	@echo "Generating sqlc code..."
	@sqlc generate

migrate-status:
	@go run cmd/migrate/main.go status

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down
