# name app
APP_NAME = server

# run dev
dev:
	go run ./cmd/${APP_NAME}
# run
run:
	docker compose up -d && go run ./cmd/${APP_NAME}
# up
up:
	docker compose up -d
# down
down:
	docker compose down

