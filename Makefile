.PHONY: wire ent generate

DB_USER=root
DB_PASSWORD=root
DB_NAME=laundry
DB_HOST=127.0.0.1
DB_PORT=5432

# Schema directory
SCHEMA_DIR := ./ent/schema

# Clean generated files but keep schema
clean:
	@echo "🧹 Cleaning generated Ent files...."
	@find ent -type f -name "*.go" ! -path "ent/schema/*" -delete
	@find ent -type d -not -path "ent/schema" -empty -delete

# Drop DB forcefully
db-drop-force:
	@echo "⚠️ Forcing drop of database $(DB_NAME)..."
	@PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -d postgres -c \
	"SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE datname='$(DB_NAME)' AND pid <> pg_backend_pid();"
	@PGPASSWORD=$(DB_PASSWORD) dropdb --if-exists -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) $(DB_NAME)
	@echo "✅ Database $(DB_NAME) dropped"

# Reset DB (drop forcefully, then recreate)
db-reset: db-drop-force
	@echo "⚠️ Dropping and recreating database $(DB_NAME)..."
	@PGPASSWORD=$(DB_PASSWORD) createdb -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) $(DB_NAME)
	@echo "✅ Database reset completed"

# Generate Wire code
wire:
	@echo "🔗 Generating Wire DI code..."
	@wire ./internal/app

# Generate Ent ORM code
ent-gen:
	@echo "⚙️  Generating Ent ORM code..."
	@ent generate ./ent/schema --feature sql/upsert 

generate: wire ent-gen
	@echo "✨ All code generation completed successfully!"

# Run the seeder
seeder:
	@echo "🌱 Running database seeders..."
	@go run ./cmd/seeder/main.go
