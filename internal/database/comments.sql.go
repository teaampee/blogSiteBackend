// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: comments.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createComment = `-- name: CreateComment :one
INSERT INTO comments (id, created_at, updated_at , content, user_id, post_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, created_at, updated_at, content, user_id, post_id
`

type CreateCommentParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Content   string
	UserID    uuid.UUID
	PostID    uuid.UUID
}

func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error) {
	row := q.db.QueryRowContext(ctx, createComment,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Content,
		arg.UserID,
		arg.PostID,
	)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Content,
		&i.UserID,
		&i.PostID,
	)
	return i, err
}

const deleteUserComment = `-- name: DeleteUserComment :exec
DELETE FROM comments where id =$1 AND user_id = $2
`

type DeleteUserCommentParams struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

func (q *Queries) DeleteUserComment(ctx context.Context, arg DeleteUserCommentParams) error {
	_, err := q.db.ExecContext(ctx, deleteUserComment, arg.ID, arg.UserID)
	return err
}

const getComment = `-- name: GetComment :one
SELECT id, created_at, updated_at, content, user_id, post_id FROM comments WHERE id = $1
`

func (q *Queries) GetComment(ctx context.Context, id uuid.UUID) (Comment, error) {
	row := q.db.QueryRowContext(ctx, getComment, id)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Content,
		&i.UserID,
		&i.PostID,
	)
	return i, err
}

const getPostComments = `-- name: GetPostComments :many
SELECT id, created_at, updated_at, content, user_id, post_id FROM comments
WHERE post_id = $3
ORDER BY created_at DESC
OFFSET $1 ROWS
LIMIT $2
`

type GetPostCommentsParams struct {
	Offset int32
	Limit  int32
	PostID uuid.UUID
}

func (q *Queries) GetPostComments(ctx context.Context, arg GetPostCommentsParams) ([]Comment, error) {
	rows, err := q.db.QueryContext(ctx, getPostComments, arg.Offset, arg.Limit, arg.PostID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Content,
			&i.UserID,
			&i.PostID,
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

const updateUserComment = `-- name: UpdateUserComment :one
UPDATE comments
SET content = $1
WHERE  id = $2 AND user_id = $3
RETURNING id, created_at, updated_at, content, user_id, post_id
`

type UpdateUserCommentParams struct {
	Content string
	ID      uuid.UUID
	UserID  uuid.UUID
}

func (q *Queries) UpdateUserComment(ctx context.Context, arg UpdateUserCommentParams) (Comment, error) {
	row := q.db.QueryRowContext(ctx, updateUserComment, arg.Content, arg.ID, arg.UserID)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Content,
		&i.UserID,
		&i.PostID,
	)
	return i, err
}
