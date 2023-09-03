postgres:
	docker run --name postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:alpine3.17

createdb: 
	docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb: 
	docker exec -it postgres dropdb simple_bank

up:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

down:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc-init:
	sqlc init

sqlc:
	sqlc generate

test: 
	go test -v -cover ./...

.PHONY: postgres createdb dropdb up down sqlc-init sqlc test