# Script: setup-sentinel.ps1
# Cria a estrutura de diretórios e arquivos para o projeto sentinel

$base = "sentinel"

# Lista de diretórios
$dirs = @(
    "$base/api",
    "$base/cmd",
    "$base/internal/config",
    "$base/internal/logging",
    "$base/internal/plugin",
    "$base/pkg/ui",
    "$base/pkg/web",
    "$base/plugins/greeter",
    "$base/webui"
)

# Lista de arquivos
$files = @(
    "$base/api/agent.proto",
    "$base/cmd/cli.go",
    "$base/cmd/gui.go",
    "$base/internal/config/config.go",
    "$base/internal/logging/logger.go",
    "$base/internal/plugin/loader.go",
    "$base/pkg/ui/manager.go",
    "$base/pkg/web/monitor.go",
    "$base/pkg/web/server.go",
    "$base/plugins/greeter/main.go",
    "$base/webui/index.html",
    "$base/main.go",
    "$base/go.mod",
    "$base/README.md"
)

Write-Host "==> Criando diretórios..."
foreach ($dir in $dirs) {
    if (-Not (Test-Path $dir)) {
        New-Item -ItemType Directory -Path $dir | Out-Null
        Write-Host "Criado: $dir"
    } else {
        Write-Host "Já existe: $dir"
    }
}

Write-Host "`n==> Criando arquivos..."
foreach ($file in $files) {
    if (-Not (Test-Path $file)) {
        New-Item -ItemType File -Path $file | Out-Null
        Write-Host "Criado: $file"
    } else {
        Write-Host "Já existe: $file"
    }
}

Write-Host "`nEstrutura concluída!"
