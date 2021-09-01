CREATE SCHEMA users;

CREATE TYPE user_status AS ENUM (
  'disabled',
  'unverified',
  'active'
);

CREATE TABLE users (
  id bigserial PRIMARY KEY,
  email varchar(254) NOT NULL UNIQUE,
  pass varchar(60) NOT NULL,
  salt varchar(60) NOT NULL,
  status user_status NOT NULL,
  verification varchar(60) NOT NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);

