*write a readme about a security tool in develepment that do many things in C/C++/C#, GO and Python*
# Palantir 🔮

**Palantir** é uma suíte modular de segurança open source.

## Componentes
- **Core (C/C++)** — gerencia plugins
- **Plugins** — port scanner, file integrity, TLS checker etc.
- **Agent API (Go)** — ingest de eventos e controle de agentes
- **Admin API (Go)** — gestão de usuários, agentes, políticas, eventos
- **Web UI (Next.js)** — painel para usuários e administradores
- **Infra** — Postgres, Redis, MinIO, Keycloak

## Quickstart
```sh
cp .env.example .env
docker-compose up -d postgres redis minio keycloak
make db-migrate
docker-compose up -d api-agent api-admin web
