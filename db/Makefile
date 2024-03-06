postgres:
	docker run -d --name my-crud -p 5432:5432 -e POSTGRES_USER=pasan -e POSTGRES_PASSWORD=12345 postgres:16-alpine

createdb:
	docker exec -it my-crud createdb --username=pasan --owner=pasan my-crud

dropdb:
	docker exec -it my-crud dropdb my-crud

migrateup:
	migrate -path db/migrations -database "postgresql://pasan:12345@localhost:5432/my-crud?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://pasan:12345@localhost:5432/my-crud?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc