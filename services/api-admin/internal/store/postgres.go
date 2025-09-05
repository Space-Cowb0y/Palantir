package store

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct{ Pool *pgxpool.Pool }

func Open(dsn string) (*Store, error) {
	p, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	return &Store{Pool: p}, nil
}
func (s *Store) Close() { s.Pool.Close() }

type Event struct {
	ID       int64          `json:"id"`
	TenantID *string        `json:"tenant_id,omitempty"`
	AgentID  *string        `json:"agent_id,omitempty"`
	Type     string         `json:"type"`
	Source   string         `json:"source"`
	Severity *string        `json:"severity,omitempty"`
	TS       time.Time      `json:"ts"`
	Payload  map[string]any `json:"payload"`
}

type EventFilter struct {
	Type, Source, Severity string
	From, To               *time.Time
	Limit, Offset          int
}

func (s *Store) QueryEvents(ctx context.Context, f EventFilter) ([]Event, int64, error) {
	args := []any{}
	where := "WHERE 1=1"
	if f.Type != "" {
		where += " AND type=$" + itoa(len(args)+1)
		args = append(args, f.Type)
	}
	if f.Source != "" {
		where += " AND source=$" + itoa(len(args)+1)
		args = append(args, f.Source)
	}
	if f.Severity != "" {
		where += " AND severity=$" + itoa(len(args)+1)
		args = append(args, f.Severity)
	}
	if f.From != nil {
		where += " AND ts >= $" + itoa(len(args)+1)
		args = append(args, *f.From)
	}
	if f.To != nil {
		where += " AND ts <  $" + itoa(len(args)+1)
		args = append(args, *f.To)
	}

	rows, err := s.Pool.Query(ctx,
		"SELECT id, tenant_id::text, agent_id::text, type, source, severity, ts, payload FROM events "+where+
			" ORDER BY ts DESC LIMIT $"+itoa(len(args)+1)+" OFFSET $"+itoa(len(args)+2),
		append(args, f.Limit, f.Offset)...,
	)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var out []Event
	for rows.Next() {
		var e Event
		if err := rows.Scan(&e.ID, &e.TenantID, &e.AgentID, &e.Type, &e.Source, &e.Severity, &e.TS, &e.Payload); err != nil {
			return nil, 0, err
		}
		out = append(out, e)
	}

	// total (sem filtros complexos; ok pro MVP)
	var total int64
	if err := s.Pool.QueryRow(ctx, "SELECT COUNT(*) FROM events "+where, args...).Scan(&total); err != nil {
		return nil, 0, err
	}
	return out, total, nil
}

func itoa(i int) string   { return fmtInt(i) }
func fmtInt(i int) string { return string('0' + i) } // simples, evita importar strconv
