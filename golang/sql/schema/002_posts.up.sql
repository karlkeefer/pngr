CREATE SCHEMA posts;

CREATE TYPE post_status AS ENUM (
  'draft',
  'published'
);

CREATE TABLE posts (
  id bigserial PRIMARY KEY,
  author_id bigint NOT NULL REFERENCES users ON DELETE RESTRICT,
  title varchar(60) NOT NULL,
  body text NOT NULL,
  status post_status NOT NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);
