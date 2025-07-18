// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: posts.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING id, created_at, updated_at, title, url, description, published_at, feed_id
`

type CreatePostParams struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Url         string
	Description string
	PublishedAt sql.NullTime
	FeedID      uuid.UUID
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Title,
		arg.Url,
		arg.Description,
		arg.PublishedAt,
		arg.FeedID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.Url,
		&i.Description,
		&i.PublishedAt,
		&i.FeedID,
	)
	return i, err
}

const deleteALLPosts = `-- name: DeleteALLPosts :exec
DELETE FROM posts
`

func (q *Queries) DeleteALLPosts(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteALLPosts)
	return err
}

const deletePost = `-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1
`

func (q *Queries) DeletePost(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deletePost, id)
	return err
}

const deletePostsByFeedID = `-- name: DeletePostsByFeedID :exec
DELETE FROM posts
WHERE feed_id = $1
`

func (q *Queries) DeletePostsByFeedID(ctx context.Context, feedID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deletePostsByFeedID, feedID)
	return err
}

const getPosts = `-- name: GetPosts :many
SELECT id, created_at, updated_at, title, url, description, published_at, feed_id
FROM posts
`

func (q *Queries) GetPosts(ctx context.Context) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Title,
			&i.Url,
			&i.Description,
			&i.PublishedAt,
			&i.FeedID,
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

const getPostsByFeedID = `-- name: GetPostsByFeedID :many
SELECT id, created_at, updated_at, title, url, description, published_at, feed_id
FROM posts
WHERE feed_id = $1
`

func (q *Queries) GetPostsByFeedID(ctx context.Context, feedID uuid.UUID) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getPostsByFeedID, feedID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Title,
			&i.Url,
			&i.Description,
			&i.PublishedAt,
			&i.FeedID,
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

const getPostsByUser = `-- name: GetPostsByUser :many
SELECT id, created_at, updated_at, title, url, description, published_at, feed_id
FROM posts
WHERE posts.feed_id IN (
    SELECT feed_follows.feed_id
    FROM feed_follows
    WHERE feed_follows.user_id = $1
)
ORDER BY posts.published_at DESC
LIMIT $2
`

type GetPostsByUserParams struct {
	UserID uuid.UUID
	Limit  int32
}

func (q *Queries) GetPostsByUser(ctx context.Context, arg GetPostsByUserParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getPostsByUser, arg.UserID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Title,
			&i.Url,
			&i.Description,
			&i.PublishedAt,
			&i.FeedID,
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

const updatePost = `-- name: UpdatePost :exec
UPDATE posts
SET updated_at = $1
WHERE id = $2
`

type UpdatePostParams struct {
	UpdatedAt time.Time
	ID        uuid.UUID
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) error {
	_, err := q.db.ExecContext(ctx, updatePost, arg.UpdatedAt, arg.ID)
	return err
}
