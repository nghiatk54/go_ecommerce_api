# name app
APP_NAME := server
# goose
GOOSE_DRIVER = "mysql"
GOOSE_DBSTRING = "root:root1234@tcp(127.0.0.1:33306)/shopDevGo"
GOOSE_MIGRATION_DIR = "sql/schema"

# command with go and docker compose
dev:
	go run ./cmd/${APP_NAME}/main.go
run:
	docker compose -f docker-compose-dev.yaml up -d &&go run ./cmd/${APP_NAME}/main.go
kill:
	docker compose -f docker-compose-dev.yaml kill
up:
	docker compose -f docker-compose-dev.yaml up -d
down:
	docker compose -f docker-compose-dev.yaml down

# command with goose
up_db:
	@GOOSE_DRIVER=${GOOSE_DRIVER} GOOSE_DBSTRING=${GOOSE_DBSTRING} goose -dir=${GOOSE_MIGRATION_DIR} up
dow_db:
	@GOOSE_DRIVER=${GOOSE_DRIVER} GOOSE_DBSTRING=${GOOSE_DBSTRING} goose -dir=${GOOSE_MIGRATION_DIR} down
reset_db:
	@GOOSE_DRIVER=${GOOSE_DRIVER} GOOSE_DBSTRING=${GOOSE_DBSTRING} goose -dir=${GOOSE_MIGRATION_DIR} reset

.PHONY: run up_db dow_db reset_db

.PHONY: air