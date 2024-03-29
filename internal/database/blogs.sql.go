// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: blogs.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createBlog = `-- name: CreateBlog :one
INSERT INTO blogs (id, created_at, updated_at , title, description, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, created_at, updated_at, title, description, user_id
`

type CreateBlogParams struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Description string
	UserID      uuid.UUID
}

func (q *Queries) CreateBlog(ctx context.Context, arg CreateBlogParams) (Blog, error) {
	row := q.db.QueryRowContext(ctx, createBlog,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Title,
		arg.Description,
		arg.UserID,
	)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.Description,
		&i.UserID,
	)
	return i, err
}

const deleteUserBlog = `-- name: DeleteUserBlog :exec
DELETE FROM blogs where user_id = $1
`

func (q *Queries) DeleteUserBlog(ctx context.Context, userID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUserBlog, userID)
	return err
}

const getBlog = `-- name: GetBlog :one
SELECT id, created_at, updated_at, title, description, user_id FROM blogs WHERE id = $1
`

func (q *Queries) GetBlog(ctx context.Context, id uuid.UUID) (Blog, error) {
	row := q.db.QueryRowContext(ctx, getBlog, id)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.Description,
		&i.UserID,
	)
	return i, err
}

const getLatestActiveBlogIDs = `-- name: GetLatestActiveBlogIDs :many
SELECT DISTINCT ON (blog_id) blog_id
FROM posts
ORDER BY created_at DESC
OFFSET $1 ROWS
LIMIT $2
`

type GetLatestActiveBlogIDsParams struct {
	Offset int32
	Limit  int32
}

func (q *Queries) GetLatestActiveBlogIDs(ctx context.Context, arg GetLatestActiveBlogIDsParams) ([]uuid.UUID, error) {
	rows, err := q.db.QueryContext(ctx, getLatestActiveBlogIDs, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []uuid.UUID
	for rows.Next() {
		var blog_id uuid.UUID
		if err := rows.Scan(&blog_id); err != nil {
			return nil, err
		}
		items = append(items, blog_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLatestBlogs = `-- name: GetLatestBlogs :many
SELECT id, created_at, updated_at, title, description, user_id FROM blogs
ORDER BY created_at DESC
OFFSET $1 ROWS
LIMIT $2
`

type GetLatestBlogsParams struct {
	Offset int32
	Limit  int32
}

func (q *Queries) GetLatestBlogs(ctx context.Context, arg GetLatestBlogsParams) ([]Blog, error) {
	rows, err := q.db.QueryContext(ctx, getLatestBlogs, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Blog
	for rows.Next() {
		var i Blog
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Title,
			&i.Description,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserBlog = `-- name: GetUserBlog :one
SELECT id, created_at, updated_at, title, description, user_id FROM blogs WHERE user_id = $1
`

func (q *Queries) GetUserBlog(ctx context.Context, userID uuid.UUID) (Blog, error) {
	row := q.db.QueryRowContext(ctx, getUserBlog, userID)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.Description,
		&i.UserID,
	)
	return i, err
}

const updateUserBlog = `-- name: UpdateUserBlog :one
UPDATE blogs
SET title = $1, description = $2
WHERE user_id = $3
RETURNING id, created_at, updated_at, title, description, user_id
`

type UpdateUserBlogParams struct {
	Title       string
	Description string
	UserID      uuid.UUID
}

func (q *Queries) UpdateUserBlog(ctx context.Context, arg UpdateUserBlogParams) (Blog, error) {
	row := q.db.QueryRowContext(ctx, updateUserBlog, arg.Title, arg.Description, arg.UserID)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.Description,
		&i.UserID,
	)
	return i, err
}
