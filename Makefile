SHELL := /bin/sh

.PHONY: dev run test fmt lint migrate-install migrate-new migrate-up migrate-down

dev:
	cd backend && air

run:
	go run ./backend/cmd/api

test:
	go test ./...

fmt:
	go fmt ./...

lint:
	golangci-lint run

migrate-install:
	go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate-new:
	migrate create -ext sql -dir backend/db/migrations -seq $(name)

migrate-up:
	migrate -path backend/db/migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path backend/db/migrations -database "$(DATABASE_URL)" down 1
