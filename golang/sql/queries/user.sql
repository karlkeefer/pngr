-- name: Create :one
INSERT INTO users (email, salt, pass, status, verification) VALUES (LOWER($1), $2, $3, $4, $5) RETURNING *;

-- name: UpdateStatus :exec
UPDATE users SET status = $1, updated_at = $2 WHERE id = $3;

-- name: FindByEmail :one
SELECT * FROM users WHERE email = LOWER($1) LIMIT 1;

-- name: FindByVerificationCode :one
SELECT * FROM users WHERE verification = $1 LIMIT 1;
