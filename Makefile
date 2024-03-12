install-swag:
	@go install github.com/swaggo/swag/cmd/swag@v1.16.3

install-sqlc:
	@go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.20.0

install-air:
	@go install github.com/cosmtrek/air@latest

install-tools: install-swag install-sqlc install-air

sqlc:
	@sqlc generate

swag-gen:
	@swag init -g ./cmd/main.go --overridesFile .swaggo

swag-fmt:
	@swag fmt

swag: swag-fmt swag-gen

generate: sqlc swag

build:
	@go build -o ./bin/main ./cmd/main.go

run: build
	@./bin/main

dev:
	@air

.PHONY: install-swag install-sqlc install-air install-tools sqlc swag-gen swag-fmt swag generate build run dev
