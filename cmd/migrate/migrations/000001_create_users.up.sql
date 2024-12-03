CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users(
	id BIGSERIAL PRIMARY KEY,
	email CITEXT UNIQUE NOT NULL, -- CITEXT is case-insensitive text
	username VARCHAR(255) UNIQUE NOT NULL,
	password BYTEA NOT NULL, -- BYTEA is binary data
	created_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW()
);

