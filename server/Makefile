postgresinit:
	docker compose up -d

postgres:
	docker exec -it postgres_db psql -U admin -d mydb

createdb:
	docker exec -it postgres_db psql -U admin -d mydb -c "CREATE DATABASE mydb;"

migration:
	migrate create -ext sql -dir database/migrations  $(name)

dropdb:
	docker exec -it postgres_db psql -U admin -d mydb -c "DROP DATABASE mydb;"

migrateup:
	migrate -path database/migrations -database "postgres://admin:secret@localhost:5432/mydb?sslmode=disable" -verbose up

migratedown:
	migrate -path database/migrations -database "postgres://admin:secret@localhost:5432/mydb?sslmode=disable" -verbose down

.PHONY: postgresinit postgres createdb createmigration dropdb migrateup migratedown