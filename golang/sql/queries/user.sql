-- name: CreateUser :one
INSERT INTO users (email, salt, pass, status, verification) VALUES (LOWER($1), $2, $3, $4, $5) RETURNING *;

-- name: UpdateUserStatus :exec
UPDATE users SET status = $2, updated_at = $3 WHERE id = $1;

-- name: FindUserByEmail :one
SELECT * FROM users WHERE email = LOWER($1) LIMIT 1;

-- name: FindUserByVerificationCode :one
SELECT * FROM users WHERE verification = $1 LIMIT 1;
