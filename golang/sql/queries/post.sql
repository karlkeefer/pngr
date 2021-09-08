-- name: CreatePost :one
INSERT INTO posts (author_id, title, body, status) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdatePost :one
UPDATE posts SET title = $3, body = $4, updated_at = NOW() WHERE id = $1 AND author_id = $2 RETURNING *;

-- name: FindPostsByAuthor :many
SELECT * FROM posts WHERE author_id = $1 ORDER BY id DESC;

-- name: FindPostByIDs :one
SELECT * FROM posts WHERE author_id = $1 AND id = $2 LIMIT 1;

-- name: DeletePostByIDs :exec
DELETE FROM posts WHERE author_id = $1 AND id = $2;
