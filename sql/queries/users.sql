-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at , name, email, password)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET name = CASE WHEN $1 <> '' THEN $1 ELSE name END,
email = CASE WHEN $2 <> '' THEN $2 ELSE email END,
password = CASE WHEN $3 <> '' THEN $3 ELSE password END
where id = $4
RETURNING *;