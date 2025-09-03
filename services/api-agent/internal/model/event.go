package model

import "time"

type Event struct {
	TS       time.Time              `json:"ts"`
	Type     string                 `json:"type"`
	Source   string                 `json:"source"`
	Severity string                 `json:"severity,omitempty"`
	AgentID  string                 `json:"agent_id"`
	Payload  map[string]any         `json:"payload"`
	TenantID string                 `json:"tenant_id,omitempty"` // opcional no ingest
}

type IngestRequest struct {
	Events []Event `json:"events"`
}
