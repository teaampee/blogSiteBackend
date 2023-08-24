-- name: CreateBlog :one
INSERT INTO blogs (id, created_at, updated_at , title, description, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetBlogs :many
SELECT * FROM blogs;

-- name: GetUserBlog :one
SELECT * FROM blogs WHERE user_id = $1;

-- name: GetBlog :one
SELECT * FROM blogs WHERE id = $1;
