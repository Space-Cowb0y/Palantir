param([Parameter(Mandatory=$true)][string]$Repo)

function Mid($title) {
  gh api "repos/$Repo/milestones" --jq ".[] | select(.title==`"$title`") | .number"
}

# Milestones
$milestones = @('M1 - Fundamentos','M2 - Plugins (3 iniciais)','M3 - Políticas','M4 - SSE & Auditoria','M5 - Hardening')
foreach ($m in $milestones) { gh api "repos/$Repo/milestones" -f title="$m" -f state='open' | Out-Null }

$M1 = Mid 'M1 - Fundamentos'; $M2 = Mid 'M2 - Plugins (3 iniciais)'; $M3 = Mid 'M3 - Políticas'; $M4 = Mid 'M4 - SSE & Auditoria'; $M5 = Mid 'M5 - Hardening'

# Labels
$labels = @('backend','frontend','core','plugin','infra','good first issue','MVP')
foreach ($l in $labels) { gh label create $l -R $Repo 2>$null }

function New-Issue($title,$body,$labels,$ms) {
  gh issue create -R $Repo --title "$title" --body "$body" --label "$labels" --milestone $ms | Out-Null
  Write-Host "✔ $title"
}

New-Issue "Agent API: persistir POST /v1/events no Postgres" @"
* Critérios:
- [ ] Validar payload
- [ ] Inserção em events
- [ ] 202 Accepted
- [ ] Teste de integração
"@ "backend,MVP" $M1

New-Issue "Admin API: listar eventos e filtro básico (GET /v1/admin/events)" @"
* Critérios:
- [ ] Paginação
- [ ] Filtros: type, source, severity, ts range
- [ ] JSON {items, total}
"@ "backend,frontend,MVP" $M1

New-Issue "Admin API: stub de SSE (GET /v1/admin/events/stream)" @"
* Critérios:
- [ ] text/event-stream
- [ ] ping periódico
"@ "backend,MVP" $M1

New-Issue "Web: tabela de eventos consumindo Admin API" @"
* Critérios:
- [ ] /events com paginação
- [ ] Filtros
- [ ] Loading/erro
"@ "frontend,MVP" $M1

New-Issue "Plugin: Port Scanner (C/C++)" "* Critérios: TCP connect, threads, eventos" "plugin,core" $M2
New-Issue "Plugin: File Integrity (C + BLAKE3)" "* Critérios: baseline, detecção, eventos" "plugin,core" $M2
New-Issue "Plugin: TLS Checker (C)" "* Critérios: versão/cadeia, eventos" "plugin,core" $M2
New-Issue "Policies: CRUD + assign + validate" "* Critérios: schema e assign" "backend,frontend" $M3
New-Issue "SSE real via LISTEN/NOTIFY" "* Critérios: NOTIFY e relay" "backend,infra" $M4
New-Issue "Hardening inicial" "* Critérios: rate-limit, CORS, headers" "backend,infra" $M5

Write-Host "🎉 Issues e milestones criados em $Repo"
