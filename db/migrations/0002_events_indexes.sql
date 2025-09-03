-- indexes
CREATE INDEX IF NOT EXISTS events_ts_tenant_idx ON events(tenant_id, ts DESC);
CREATE INDEX IF NOT EXISTS events_payload_gin ON events USING gin (payload);
