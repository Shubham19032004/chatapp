# Load environment variables from .env
SHELL := /bin/bash
APP_ENV := ./backend/app.env

include $(APP_ENV)
export $(shell sed 's/=.*//' $(APP_ENV))
postgres:
	docker run --name  postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root chatapp

dropdb:
	docker exec -it postgres dropdb chatapp
# Migrate up
migrateup:
	@source $(APP_ENV) && migrate -path backend/db/migration -database "$$DB_URL" -verbose up

# Migrate up 1 step
migrateup1:
	@source $(APP_ENV) && migrate -path backend/db/migration -database "$$DB_URL" -verbose up 1

# Migrate down
migratedown:
	@source $(APP_ENV) && migrate -path backend/db/migration -database "$$DB_URL" -verbose down

# Migrate down 1 step
migratedown1:
	@source $(APP_ENV) && migrate -path backend/db/migration -database "$$DB_URL" -verbose down 1

# Create migration
migration:
	migrate create -ext sql -dir backend/db/migration -seq init_schema

# Generate SQLC code
sqlc:
	sqlc generate

# Start the server
server:
	go run backend/main.go

# Run tests
test:
	go test -v -cover ./...

# Generate mocks using mockgen
mock:
	mockgen -package mockdb -destination backend/db/mock/store.go bank/db/sqlc

# Build database documentation
db_docs:
	dbdocs build docs/db.dbml

# Generate SQL schema from DBML
db_schema:
	dbml2sql --postgres -o docs/schema.sql docs/db.dbml


proto:
	rm -f backend/pb/*.go
	protoc --proto_path=backend/proto --go_out=backend/pb --go_opt=paths=source_relative \
	    --go-grpc_out=backend/pb --go-grpc_opt=paths=source_relative \
	    --grpc-gateway_out=backend/pb --grpc-gateway_opt=paths=source_relative \
	    backend/proto/*.proto
check:
	@source .env && echo $$DB_URL

.PHONY: migration migrateup migratedown migrateup1 migratedown1 sqlc test server mock db_docs db_schema proto check
