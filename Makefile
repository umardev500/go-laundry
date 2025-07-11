.PHONY: ent-clean ent wire seed reset-db

# Load variables from .env file
ifneq (,$(wildcard .env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

seed:
	@go run cmd/seed/main.go 

ent-clean:
	@find ./internal/ent -mindepth 1 -maxdepth 1 ! -name schema -exec rm -rf {} +

ent:
	ent generate --feature sql/upsert ./internal/ent/schema

reset-db:
	@echo "Dropping all tables in $(DB_NAME)..."
	PGPASSWORD=$(DB_PASS) psql -h $(DB_HOST) -U $(DB_USER) -d $(DB_NAME) -c "DROP SCHEMA public CASCADE; CREATE SCHEMA public;"

wire:
	@wire ./internal/di