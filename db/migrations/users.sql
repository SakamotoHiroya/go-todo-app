CREATE TYPE oauth_provider AS ENUM (
  'google'
);

CREATE TYPE user_status AS ENUM (
  'active',
  'inactive',
  'suspended'
);

CREATE TABLE users (
  id                  BIGSERIAL PRIMARY KEY,
  provider            oauth_provider NOT NULL DEFAULT 'google',
  provider_user_id    VARCHAR(128)    NOT NULL,
  email               VARCHAR(255)    NOT NULL,
  name                VARCHAR(100),
  avatar_url          VARCHAR(512),
  status              user_status     NOT NULL DEFAULT 'active',
  created_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
  updated_at          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),

  CONSTRAINT uq_provider_provider_user
    UNIQUE (provider, provider_user_id),
  CONSTRAINT uq_users_email
    UNIQUE (email)
);

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_users_updated_at
  BEFORE UPDATE ON users
  FOR EACH ROW
  EXECUTE PROCEDURE trigger_set_timestamp();