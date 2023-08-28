migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/avito_backend?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/avito_backend?sslmode=disable" -verbose down
generate:
	sqlc generate
.PHONY: migrateup, migratedown, generate
