-- name: CreateBlogPost :one
INSERT INTO posts (id, created_at, updated_at , title, content, user_id, blog_id)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetBlogPosts :many
SELECT * FROM posts
WHERE blog_id = $1
ORDER BY created_at DESC
LIMIT 50;

-- name: GetPost :one
SELECT * FROM posts WHERE id = $1;

-- name: UpdateUserPost :one
UPDATE posts
SET title = $1, content = $2
WHERE  id = $3 AND user_id = $4
RETURNING *;

-- name: DeleteUserPost :exec
DELETE FROM posts where id =$1 AND user_id = $2;
