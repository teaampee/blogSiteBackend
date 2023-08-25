-- name: CreateBlog :one
INSERT INTO blogs (id, created_at, updated_at , title, description, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetBlogs :many
SELECT * FROM blogs
ORDER BY updated_at ASC
LIMIT 50;

-- name: GetUserBlog :one
SELECT * FROM blogs WHERE user_id = $1;

-- name: GetBlog :one
SELECT * FROM blogs WHERE id = $1;

-- name: UpdateUserBlog :one
UPDATE blogs
SET title = $1, description = $2
WHERE user_id = $3
RETURNING *;

-- name: DeleteUserBlog :exec
DELETE FROM blogs where user_id = $1;
