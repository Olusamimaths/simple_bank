postgres:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=${POSTGRES_USER} -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -d postgres:17-alpine

createdb:
	docker exec -it postgres17 createdb --username=${POSTGRES_USER} --owner=${POSTGRES_USER} simple_bank

dropdb:
	docker exec -it postgres17 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "${DB_SOURCE}" -verbose up

migrateup_latest:
	migrate -path db/migration -database "${DB_SOURCE}" -verbose up 1

migratedown:
	migrate -path db/migration -database "${DB_SOURCE}" -verbose down

migratedown_latest:
	migrate -path db/migration -database "${DB_SOURCE}" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/Olusamimaths/simple_bank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migrateup_latest migratedown migratedown_latest sqlc test server mock