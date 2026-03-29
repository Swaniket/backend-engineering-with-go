CREATE EXTENSION IF NOT EXISTS citext; -- Remove this if we do this is db_init.sql

CREATE TABLE IF NOT EXISTS users(
    id bigserial PRIMARY KEY,
    email citext UNIQUE NOT NULL, -- citext is an extension that we use for case-insensitive text storage and comparison 
    username varchar(255) UNIQUE NOT NULL,
    password bytea NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
)