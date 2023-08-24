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

const getBlogs = `-- name: GetBlogs :many
SELECT id, created_at, updated_at, title, description, user_id FROM blogs
`

func (q *Queries) GetBlogs(ctx context.Context) ([]Blog, error) {
	rows, err := q.db.QueryContext(ctx, getBlogs)
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
