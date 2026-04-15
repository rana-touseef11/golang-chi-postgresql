include .env
export

print:
	echo $(DB_URL)

run-dev:
	go run cmd/api/main.go

build:
	go build -o app ./cmd/api

start:
	./app

swag-init:
	swag init -g cmd/api/main.go

# DB cmd
create-db:
	createdb -U $(DB_USER) -h $(DB_HOST) -p $(DB_PORT) $(DB_NAME) || true

migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

migrate-up:
	migrate -path migrations -database $(DB_URL) up

migrate-down:
	migrate -path migrations -database $(DB_URL) down
