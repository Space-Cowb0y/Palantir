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

# To-Do List

- [X] Definir requisitos do projeto
- [X] Criar estrutura de pastas
- [ ] Implementar funcionalidades principais
- [ ] Testar aplicação
- [ ] Documentar código
- [ ] Revisar e refatorar
- [ ] Publicar projeto
- [ ] -----------------------------------------
- [ ] Plugins:
  - [ ] Scanner de portas
  - [ ] Verificador de integridade de arquivos
  - [ ] Verificador de TLS
  - [ ] Cofre de Senhas
  - [ ] Honeypot
  - [ ] Mini NIDS
  - [ ] DNS sinkhole local
  - [ ] Secrets scanner
  - [ ] Sandbox de execução
  - [ ] Policy enforcer
  - [ ] Format parser com propriedade
  - [ ] Fuzzer orientado a cobertura
  - [ ] Assinaturas e verificação
  - [ ] Encrypted backup deduplicado
  - [ ] Key management
  - [ ] Key management
  - [ ] OAuth2/OIDC client e resource server
  - [ ] SCA/Dependency auditor
- [] C/C++
  - [ ] ELF parser + verificador de RPATH inseguro
  - [ ] Implementação de Safe string libs
  - [ ] Mini-TLS (client)
  - [ ] Kernel driver toy
- [] Python
  - [ ] YARA runner
  - [ ] Threat intel aggregator
  - [ ] Forense rápida de navegador
- [] RUST
  - [ ] BLAKE3-based file hasher multi-thread
  - [ ] PCAP sniffer zero-copy
  - [ ] WASM sandbox runner
  - [ ] CLI de criptografia com zeroize
- [] GO
  - [ ] Agent de endpoint
  - [ ] Agent de endpoint
  - [ ] Scanner de configuração cloud
  - [ ] Rate-limiter distribuído
