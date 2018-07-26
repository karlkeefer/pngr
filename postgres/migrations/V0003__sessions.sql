CREATE SCHEMA sessions;

CREATE TABLE sessions (
	sid			VARCHAR(128)	PRIMARY KEY,
	ip			VARCHAR(128)	NOT NULL,
	user_agent	VARCHAR(128)	NOT NULL,
	created		TIMESTAMP 		NOT NULL DEFAULT now(),
	expires		TIMESTAMP
);

INSERT INTO ops (op) VALUES('migration V0003__sessions.sql');