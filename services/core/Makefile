include .env

local_db_connection_url ?= postgresql://postgres:postgres@localhost:5432/scrumerization?sslmode=disable
migration_path ?= internal/infrastructure/db/postgres/
step ?= 1

dev:
	go run cmd/scrumerization/main.go
gql-generate:
	go run github.com/99designs/gqlgen generate
migrate-create:
	migrate create -ext sql -dir ${migration_path} -seq $(name)
migrate-up:
	migrate -path ${migration_path} -database "$(DB_SOURCE)" -verbose up
migrate-up-ci:
	migrate -path ${migration_path} -database "${local_db_connection_url}" -verbose up
migrate-down-ci:
	migrate -path ${migration_path} -database "$(local_db_connection_url)" -verbose down
migrate-reset:
	migrate -path ${migration_path} -database "$(local_db_connection_url)" -verbose drop
migrate-down:
	migrate -path ${migration_path} -database "$(DB_SOURCE)" -verbose down
migrate-up-step:
    ifndef step
        $(error step is not set. Please provide a value representing number.)
    endif
	migrate -path ${migration_path} -database "$(DB_SOURCE)" -verbose up $(step)
migrate-down-step:
    ifndef step
        $(error step is not set. Please provide a value representing number.)
    endif
	migrate -path ${migration_path} -database "$(DB_SOURCE)" -verbose down $(step)

.PHONY: dev gql-generate migrate-create migrate-up migrate-up-ci migrate-down migrate-up-step migrate-up-ci migrate-down-ci migrate-down-step migrate-reset