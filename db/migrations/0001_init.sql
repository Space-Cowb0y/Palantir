CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE TABLE IF NOT EXISTS tenants(
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  name text UNIQUE NOT NULL,
  created_at timestamptz DEFAULT now()
);

CREATE TABLE IF NOT EXISTS users(
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  tenant_id uuid REFERENCES tenants(id),
  email text UNIQUE NOT NULL,
  name text NOT NULL,
  password_hash text,
  oidc_sub text,
  disabled bool DEFAULT false,
  created_at timestamptz DEFAULT now()
);

CREATE TABLE IF NOT EXISTS agents(
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  tenant_id uuid REFERENCES tenants(id),
  name text NOT NULL,
  node_id text UNIQUE,
  version text,
  os text, arch text,
  tags jsonb DEFAULT '{}'::jsonb,
  last_seen timestamptz
);

CREATE TABLE IF NOT EXISTS events(
  id bigserial PRIMARY KEY,
  tenant_id uuid REFERENCES tenants(id),
  agent_id uuid REFERENCES agents(id),
  source text, type text, severity text,
  ts timestamptz NOT NULL,
  payload jsonb NOT NULL
);
