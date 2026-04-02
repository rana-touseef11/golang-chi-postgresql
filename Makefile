include .env
export

print:
	echo $(DB_URL)

run-start:
	go run cmd/api/main.go

swag-init:
	swag init -g cmd/api/main.go

migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

migrate-up:
	migrate -path migrations -database $(DB_URL) up

migrate-down:
	migrate -path migrations -database $(DB_URL) down
