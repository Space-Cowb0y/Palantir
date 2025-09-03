.PHONY: all core plugins apis web db-migrate

all: core plugins apis web

core:
	@cmake -S . -B build && cmake --build build --target core

plugins:
	@cmake --build build --target port_scanner file_integrity tls_checker

apis:
	@(cd services/api-agent && go build ./cmd/api-agent)
	@(cd services/api-admin && go build ./cmd/api-admin)

web:
	@(cd web && npm install && npm run dev)

db-migrate:
	@psql $(DATABASE_URL) -f db/migrations/0001_init.sql && \
	 psql $(DATABASE_URL) -f db/migrations/0002_events_indexes.sql && \
	 psql $(DATABASE_URL) -f db/migrations/0003_seed_roles.sql
