-- name: CreateComment :one
INSERT INTO comments (id, created_at, updated_at , content, user_id, post_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetPostComments :many
SELECT * FROM comments
WHERE post_id = $3
ORDER BY created_at DESC
OFFSET $1 ROWS
LIMIT $2;

-- name: GetComment :one
SELECT * FROM comments WHERE id = $1;

-- name: UpdateUserComment :one
UPDATE comments
SET content = $1
WHERE  id = $2 AND user_id = $3
RETURNING *;

-- name: DeleteUserComment :exec
DELETE FROM comments where id =$1 AND user_id = $2;
