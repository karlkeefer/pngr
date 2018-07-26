CREATE SCHEMA ops;

CREATE TABLE ops (
	id			BIGSERIAL		PRIMARY KEY,
	op			VARCHAR(254)	UNIQUE
);

INSERT INTO ops (op) VALUES('migration V0001__migrations.sql');