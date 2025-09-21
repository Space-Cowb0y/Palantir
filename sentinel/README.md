# Sentinel (control plane)

Sentinel gerencia **Eyes** (plugins externos) via gRPC. Este repositório traz CLI (Cobra), TUI (Bubble Tea) e serviços base (gRPC + HTTP) prontos para evoluir.

## Requisitos
- Go 1.25.1
- (Opcional) protoc para gerar stubs gRPC

## Rodando
```bash
cd sentinel
go run . run
# Em outro terminal
go run . ui