REQUIREMENTS — Palantir v0.1 (MVP)
1) Visão & Objetivos

Palantir é uma suíte modular de segurança com núcleo em C/C++ (plugins), duas APIs em Go (Agent/Admin) e UI web (Next.js).

Objetivo do MVP (v0.1): coletar eventos de 3 plugins iniciais, exibir no painel, aplicar políticas simples e gerenciar usuários/agents.

2) Escopo (MVP vs. Pós-MVP)

MVP (v0.1)

Core C/C++ com carregamento de plugins (processo isolado), assinatura opcional (stub).

Plugins iniciais:

Scanner de portas (C/C++)

Verificador de integridade de arquivos (C com BLAKE3)

Verificador de TLS (C com mbedTLS/wolfSSL)

Agent API (Go): POST /v1/events, GET /v1/policies, heartbeats, upload básico de artefatos (MinIO).

Admin API (Go): CRUD básico de usuários/agents/policies, SSE /events/stream.

UI: login OIDC (Keycloak), dashboards (eventos, agentes), listagem/assign de policies.

Banco: Postgres (particionamento poderá vir depois), Redis (fila simples opcional), MinIO (artefatos).

AuthN/AuthZ: OIDC (Keycloak) + RBAC (admin/analyst/viewer).

Pós-MVP (v0.2+)

Honeypot, Mini NIDS, DNS sinkhole.

Secrets scanner, Sandbox de execução (seccomp/namespaces), Policy enforcer.

SCA auditor, OAuth2/OIDC client/resource server (para testes externos).

Análise forense, Threat intel, etc. (demais itens da sua lista).

3) Personas & Casos de Uso

Admin: cria usuários/roles, registra/gerencia agents, define e atribui policies.

Analyst: acompanha eventos em tempo real, consulta histórico, exporta artefatos.

Operator: rola/rollback de policies, executa comandos em agents.

Agent: publica eventos e recebe políticas/comandos.

Casos (MVP)

Como Admin, quero criar uma policy “TLS mínimo 1.2” e atribuir a um agente.

Como Agent, quero enviar eventos do plugin Port Scanner para o backend.

Como Analyst, quero ver em tempo real eventos por severidade e aplicar filtros.

4) Requisitos Funcionais

RF-01: Core deve carregar/descarregar plugins com ABI C estável.

RF-02: Plugins executam em processo isolado por padrão (spawn) e se comunicam com o Core.

RF-03: Core publica eventos na Agent API (HTTP).

RF-04: Agent API persiste eventos no Postgres e (opcional) em fila.

RF-05: Admin API expõe CRUD de usuários/agents/policies e stream SSE de eventos.

RF-06: UI autentica via OIDC e consome somente Admin API.

RF-07: Policies são consultadas pelos Agents (pull) e exibidas/gerenciadas no Admin.

5) Requisitos Não Funcionais

RNF-01: Disponibilidade alvo MVP: 99% (lab/dev).

RNF-02: Desempenho ingest MVP: 2k eventos/segundo em máquina única (escala horizontal futura).

RNF-03: Segurança: TLS em produção, JWT válido, rate-limit básico, logs de auditoria.

RNF-04: Portabilidade: Linux x86_64 (prioridade), Windows suporte básico (loader adaptado).

RNF-05: Observabilidade: logs estruturados, métricas Prometheus, traces (OpenTelemetry) nas APIs.

6) Arquitetura (resumo)

Core (C/C++) ↔ Plugins (processos)

Core → Agent API (Go): HTTP ingest/artefatos

Admin API (Go): gestão/consultas/stream

UI (Next.js) ↔ Admin API

Postgres (dados), MinIO (artefatos), Redis (fila opcional), Keycloak (OIDC)

7) Modelo de Dados (MVP)

tenants(id, name, created_at)

users(id, tenant_id, email, name, disabled, created_at)

agents(id, tenant_id, name, node_id, version, os, arch, tags, last_seen)

events(id, tenant_id, agent_id, source, type, severity, ts, payload jsonb)

policies(id, tenant_id, name, type, version, status, spec jsonb, created_at)

policy_assignments(policy_id, agent_id, assigned_by, assigned_at, status)

audit_logs(id, tenant_id, actor_user_id, action, resource, ts, meta)

8) APIs (MVP)

Agent API

POST /v1/events — batch de eventos (202 Accepted)

POST /v1/heartbeats — estado do agent/plugins

GET /v1/policies?agent_id= — políticas aplicáveis

POST /v1/artifacts — upload p/ MinIO

Admin API

GET /v1/admin/users — listar (viewer+)

GET /v1/admin/agents — listar

GET/POST /v1/admin/policies — CRUD

POST /v1/admin/policies/{id}/assign — assign a agents

GET /v1/admin/events/stream — SSE (viewer+)

(Demais CRUDs podem ser adicionados progressivamente)

9) Segurança

OIDC (Keycloak), tokens com roles (admin/analyst/viewer).

Rate-limit na Agent/Admin API.

Logs de auditoria em operações sensíveis.

Em produção: TLS, cabeçalhos seguros, cookies SameSite, CORS restrito.

Plugins rodando com least privilege; futura sandbox seccomp/namespaces.

10) Deploy & Ambientes

Local: docker-compose (Postgres, Redis, MinIO, Keycloak, APIs, Web).

Prod (futuro): Kubernetes (manifests em deploy/k8s/), imagens com distroless/ubi.

11) Observabilidade

Logs JSON nas APIs; métricas (requests, latência, ingest TPS).

Painéis (Grafana) pós-MVP.

12) Testes

Unit: Core loader, parsers, handlers HTTP.

Integração: Postgres + Agent/Admin API; ingest end-to-end.

E2E: subir stack local, gerar eventos sintéticos, validar UI.

Segurança: validação de JWT/JWKS e RBAC.

13) Documentação

README.md (quickstart), docs/REQUIREMENTS.md, docs/architecture.md, docs/openapi/*.yaml.

Comentários em código e exemplos de chamadas (curl).

14) Critérios de Aceite (MVP)

 Subir stack via docker-compose e logar na UI via OIDC.

 Criar uma policy no Admin, assignar a um agent.

 Enviar lote de eventos (scanner/integ/TLS) pela Agent API e visualizar na UI (lista/contagem).

 SSE entrega eventos novos em tempo quase real (<3s).

 RBAC efetivo (viewer não altera policies).

 Testes unitários básicos passam no CI.

Roadmap do Item 2: Implementar funcionalidades principais (MVP)

Milestone M1 — Fundamentos

 Persistência real no Agent API (POST /v1/events → Postgres).

 Admin API: rotas users, agents, policies, events/stream (SSE dummy → real depois).

 UI: login OIDC + dashboard com contadores e tabela de eventos.

Milestone M2 — Plugins (3 primeiros)

 Port Scanner: threads, TCP connect, banner opcional → envia eventos.

 File Integrity: BLAKE3 + baseline, detecta alteração → evento.

 TLS Checker: coleta cadeia, validade, versão mínima → evento/severidade.

Milestone M3 — Políticas

 Estrutura policies.spec (JSON Schema simples).

 Admin: CRUD + assign; Agent: GET /v1/policies e cache local.

 Enforcements simples (ex.: TLS <1.2 = severity “high”).

Milestone M4 — SSE real e Auditoria

 LISTEN/NOTIFY no Postgres ou fila Redis para /events/stream.

 audit_logs para operações administrativas.

Milestone M5 — Hardening básico

 Rate-limit, CORS restrito, headers seguros.

 Execução de plugins em processo com flags de “no new privs”.