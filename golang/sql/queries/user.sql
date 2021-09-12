-- name: CreateUser :one
INSERT INTO users (email, salt, pass, status, verification)
  VALUES (LOWER(@email::varchar), @salt::varchar, @pass::varchar, @status::user_status, @verification::varchar) RETURNING *;

-- name: UpdateUserStatus :exec
UPDATE users SET status = $2, updated_at = NOW() WHERE id = $1;

-- name: UpdateUserPassword :exec
UPDATE users SET salt = $2, pass = $3, updated_at = NOW() WHERE id = $1;

-- name: FindUserByID :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: FindUserByEmail :one
SELECT * FROM users WHERE email = LOWER($1) LIMIT 1;

-- name: FindUserByVerificationCode :one
SELECT * FROM users WHERE verification = $1 LIMIT 1;
