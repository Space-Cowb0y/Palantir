.PHONY: all core plugins apis web db-migrate db-shell

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

# -----------------------
# DB (via container)
# -----------------------

# nome do serviço no docker-compose
DB_SVC  ?= postgres
DB_USER ?= security
DB_NAME ?= security
# Se quiser forçar usuário/db por variável, defina DB_USER/DB_NAME ao chamar o make.
# Caso contrário, usamos as variáveis *do container* ($$POSTGRES_USER/$$POSTGRES_DB).

# Macro que executa psql dentro do container lendo SQL do stdin (-f -)
PSQL_IN_CONTAINER = docker compose exec -T $(DB_SVC) psql -v ON_ERROR_STOP=1 -U $(DB_USER) -d $(DB_NAME) -f -


db-migrate:
	@echo ">> applying 0001_init.sql"
	@$(PSQL_IN_CONTAINER) < db/migrations/0001_init.sql
	@echo ">> applying 0002_events_indexes.sql"
	@$(PSQL_IN_CONTAINER) < db/migrations/0002_events_indexes.sql
	@echo ">> applying 0003_seed_roles.sql"
	@$(PSQL_IN_CONTAINER) < db/migrations/0003_seed_roles.sql
	@echo ">> migrations applied."

# abre um shell psql interativo dentro do container (útil pra inspeção)
db-shell:
	@docker compose exec -it $(DB_SVC) sh -lc 'psql -U $${POSTGRES_USER:-$(DB_USER)} -d $${POSTGRES_DB:-$(DB_NAME)}'

db-migrate-local:
	@psql $(DATABASE_URL) -f db/migrations/0001_init.sql && \
	 psql $(DATABASE_URL) -f db/migrations/0002_events_indexes.sql && \
	 psql $(DATABASE_URL) -f db/migrations/0003_seed_roles.sql
