-- name: CreatePost :one
INSERT INTO posts (author_id, title, body, status) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdatePost :one
UPDATE posts SET title = $1, body = $2, updated_at = $3 WHERE id = $4 AND author_id = $5 RETURNING *;

-- name: FindPostsByAuthor :many
SELECT * FROM posts WHERE author_id = $1 ORDER BY id DESC;

-- name: FindPostByIDs :one
SELECT * FROM posts WHERE author_id = $1 AND id = $2 LIMIT 1;

-- name: DeletePostByIDs :exec
DELETE FROM posts WHERE author_id = $1 AND id = $2;
