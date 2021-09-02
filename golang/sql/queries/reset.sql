-- name: CreateReset :one
INSERT INTO resets (user_id, code) VALUES ($1, $2) RETURNING *;

-- name: FindResetByCode :one
SELECT * FROM resets WHERE code = $1 LIMIT 1;

-- name: DeleteResetsForUser :exec
DELETE FROM resets WHERE user_id = $1;
