package store

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct{ Pool *pgxpool.Pool }

func Open(dsn string) (*Store, error) {
	p, err := pgxpool.New(context.Background(), dsn)
	if err != nil { return nil, err }
	return &Store{Pool: p}, nil
}
func (s *Store) Close(){ s.Pool.Close() }

type EventRow struct {
	TenantID *string
	AgentID  *string
	Type     string
	Source   string
	Severity *string
	TS       time.Time
	Payload  map[string]any
}

func (s *Store) InsertEvents(ctx context.Context, rows []EventRow) error {
	batch := &pgx.Batch{}
	for _, r := range rows {
		batch.Queue(
			`INSERT INTO events(tenant_id, agent_id, type, source, severity, ts, payload)
			 VALUES ($1,$2,$3,$4,$5,$6,$7)`,
			r.TenantID, r.AgentID, r.Type, r.Source, r.Severity, r.TS, r.Payload,
		)
	}
	br := s.Pool.SendBatch(ctx, batch)
	defer br.Close()
	for range rows {
		if _, err := br.Exec(); err != nil { return err }
	}
	return nil
}
