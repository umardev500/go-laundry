# Makefile

DB_USER=root
DB_PASSWORD=root
DB_NAME=laundry
DB_HOST=127.0.0.1
DB_PORT=5432

# Schema directory
SCHEMA_DIR := ./ent/schema

# -------------------------------
# Commands
# -------------------------------

# Create a new schema
new:
	@if [ -z "$(name)" ]; then \
		echo "Please provide schema name: make new name=Pet"; \
		exit 1; \
	fi
	@echo "Creating new Ent schema: $(name)"
	@go run -mod=mod entgo.io/ent/cmd/ent new $(name)
	@snake_case=$$(echo $(name) | sed -E 's/([a-z0-9])([A-Z])/\1_\L\2/g' | tr '[:upper:]' '[:lower:]'); \
	if [ -f "ent/schema/$${snake_case}.go" ]; then \
		echo "Schema already exists: ent/schema/$${snake_case}.go"; \
	else \
		mv ent/schema/$$(echo $(name) | tr '[:upper:]' '[:lower:]').go ent/schema/$${snake_case}.go; \
	fi

# Generate code from schema
generate:
	@echo "⚙️  Generating Ent code..."
	@go run -mod=mod entgo.io/ent/cmd/ent generate $(SCHEMA_DIR) --feature sql/upsert

# Run migrations (optional)
migrate:
	@echo "📦 Running Ent migrations..."
	@go run ./cmd/migrate/migrate.go

# Clean generated files but keep schema
clean:
	@echo "🧹 Cleaning generated Ent files...."
	@find ent -type f -name "*.go" ! -path "ent/schema/*" -delete
	@find ent -type d -not -path "ent/schema" -empty -delete

# Generate Wire code
wire:
	@echo "🔗 Generating Wire DI code..."
	@wire ./internal/app

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

# Reset DB and apply migrations
db-recreate: db-reset migrate

# Seed database with initial data
db-seed:
	@echo "🌱 Seeding database..."
	go run cmd/seed/main.go

# Help
help:
	@echo "Makefile commands:"
	@echo "  make new name=User       # Create new Ent schema"
	@echo "  make generate            # Generate Ent code from schema"
	@echo "  make migrate             # Run migrations"
	@echo "  make db-drop-force       # Drop database"
	@echo "  make db-reset            # Drop & recreate database"
	@echo "  make db-recreate         # Drop, recreate, and run migrations"
	@echo "  make db-seed             # Seed database with initial data"
	@echo "  make clean               # Remove generated files"
	@echo "  make wire                # Generate Wire DI code"
