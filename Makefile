postgres:
	docker run --name postgres_pizza -p 5434:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.4-alpine  

createdb:
	docker exec -it postgres_pizza createdb --username=root --owner=root pizza_order

dropdb:
	docker exec -it postgres_pizza dropdb pizza_order

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5434/pizza_order?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5434/pizza_order?sslmode=disable" -verbose down

sqlc:
	docker run --rm -v "$(CURDIR):/src" -w //src kjconroy/sqlc generate

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server
