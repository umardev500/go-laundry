.PHONY: create_structure create_schema generate_ent ent-clean init-schema wire seed

# Load variables from .env file
ifneq (,$(wildcard .env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

seed:
	@go run cmd/seed/main.go 

init-schema:
	@echo "Initializing schema $(name)"
	@ent new $(name) --target ./internal/ent/schema

ent-clean:
	@find ./internal/ent -mindepth 1 -maxdepth 1 ! -name schema -exec rm -rf {} +

gen-ent:
	ent generate --feature sql/upsert ./internal/ent/schema

gen-schema:
ifndef name
	$(error name is required, usage: make create_schema name=User)
endif
	./script/create-schema.sh $(name)

gen-structure:
ifndef name
	$(error name is required, usage: make create_structure name=DomainName)
endif
	./script/create-structure.sh $(name)

drop-all:
	@echo "Dropping all tables in $(DB_NAME)..."
	PGPASSWORD=$(DB_PASS) psql -h $(DB_HOST) -U $(DB_USER) -d $(DB_NAME) -c "DROP SCHEMA public CASCADE; CREATE SCHEMA public;"

wire:
	@wire ./internal/di