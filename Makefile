.PHONY: build
build:
	go build -v ./cmd/server

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: migrate
migrate:
	migrate -path ./migrations -database 'postgres://atcer:sunsetdivine@localhost:5432/atc?sslmode=disable' up

.DEFAULT_GOAL := build