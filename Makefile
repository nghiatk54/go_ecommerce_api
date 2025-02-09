# name app
APP_NAME := server
# goose
GOOSE_DRIVER = "mysql"
GOOSE_DBSTRING = "root:root1234@tcp(127.0.0.1:33306)/shopDevGo"
GOOSE_MIGRATION_DIR = "sql/schema"

# command with go and docker compose
docker_build:
	docker compose -f docker-compose-dev.yaml up -d --build
	docker compose ps
docker_up:
	docker compose -f docker-compose-dev.yaml up -d
docker_stop:
	docker compose -f docker-compose-dev.yaml down
dev:
	go run ./cmd/${APP_NAME}/main.go
run:
	docker compose -f docker-compose-dev.yaml up -d && go run ./cmd/${APP_NAME}/main.go

# command with goose
up_by_one_db:
	@GOOSE_DRIVER=${GOOSE_DRIVER} GOOSE_DBSTRING=${GOOSE_DBSTRING} goose -dir=${GOOSE_MIGRATION_DIR} up-by-one
up_all_db:
	@GOOSE_DRIVER=${GOOSE_DRIVER} GOOSE_DBSTRING=${GOOSE_DBSTRING} goose -dir=${GOOSE_MIGRATION_DIR} up
down_all_db:
	@GOOSE_DRIVER=${GOOSE_DRIVER} GOOSE_DBSTRING=${GOOSE_DBSTRING} goose -dir=${GOOSE_MIGRATION_DIR} down
reset_all_db:
	@GOOSE_DRIVER=${GOOSE_DRIVER} GOOSE_DBSTRING=${GOOSE_DBSTRING} goose -dir=${GOOSE_MIGRATION_DIR} reset
create_migration:
	@goose -dir=${GOOSE_MIGRATION_DIR} create ${name} sql
# command with sqlc
sqlc_generate:
	sqlc generate
# command with air
air:
	air
#command with swagger
swagger:
	swag init -g ./cmd/${APP_NAME}/main.go -o ./cmd/swagger/docs


.PHONY: dev run docker_up docker_stop docker_build
.PHONY: up_by_one_db up_all_db down_all_db reset_all_db create_migration
.PHONY: sqlc_generate 
.PHONY: air
.PHONY: swagger