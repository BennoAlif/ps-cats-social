GO := go
MIGRATE := migrate

# Load environment variables and construct DB_DSN  
include .env
export

DB_DSN := postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

.PHONY: tidy run build test db-up db-down migrate-up migrate-down migrate-new
tidy:  		; $(GO) mod tidy
build: 		; $(GO) build -o bin/api ./src
run-build: build; DB_USERNAME=$(DB_USER) DB_PASSWORD=$(DB_PASSWORD) DB_HOST=$(DB_HOST) DB_PORT=$(DB_PORT) DB_NAME=$(DB_NAME) DB_PARAMS=sslmode=disable ./bin/api --http-addr=:8080
run:   		; DB_USERNAME=$(DB_USER) DB_PASSWORD=$(DB_PASSWORD) DB_HOST=$(DB_HOST) DB_PORT=$(DB_PORT) DB_NAME=$(DB_NAME) DB_PARAMS=sslmode=disable $(GO) run ./src --http-addr=:8080
test:  		; $(GO) test ./... -race -count=1
migrate-up:  ; $(MIGRATE) -path db/migrations -database "$(DB_DSN)" up
migrate-down:; $(MIGRATE) -path db/migrations -database "$(DB_DSN)" down 1
migrate-new: ; @test -n "$(name)"; $(MIGRATE) create -ext sql -dir db/migrations -seq $(name)
db-up: 		; docker compose up -d postgres
db-down: 	; docker compose down -v

