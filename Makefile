postgres:
	docker run --name postgres16 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root  assistant_teacher

dropdb:
	docker exec -it postgres16 dropdb assistant_teacher

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/assistant_teacher?sslmode=disable" -verbose up 

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/assistant_teacher?sslmode=disable" -verbose down 

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test