CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users(
	id bigserial PRIMARY KEY,
	email citext UNIQUE NOT NULL, -- citext is case-insensitive text
	username varchar(255) UNIQUE NOT NULL,
	password bytea NOT NULL, -- bytea is binary data
	created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);
