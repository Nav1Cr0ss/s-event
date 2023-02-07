SERVICE=s-event
DB_URL=postgresql://nav1cr0ss:0608@localhost:5436/$(SERVICE)?sslmode=disable

postgres:
	docker run --name postgres-NearBy -p 5436:5432 -e POSTGRES_USER=nav1cr0ss -e POSTGRES_PASSWORD=0608 -d postgres:14-alpine

createdb:
	docker exec -it postgres-NearBy createdb --username=nav1cr0ss $(SERVICE)

dropdb:
	docker exec -it postgres-NearBy dropdb --username=nav1cr0ss "$(SERVICE)"

migrateup:
	migrate -path internal/adapters/repository/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path internal/adapters/repository/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc --file internal/adapters/repository/sqlcConfig.yaml generate




.PHONY: sqlc migratedown migrateup dropdb createdb postgres