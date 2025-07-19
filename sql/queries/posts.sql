-- name: CreatePost :one
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
RETURNING *;

-- name: GetPosts :many
SELECT *
FROM posts;

-- name: DeleteALLPosts :exec
DELETE FROM posts;

-- name: GetPostsByFeedID :many
SELECT *
FROM posts
WHERE feed_id = $1;

-- name: DeletePostsByFeedID :exec
DELETE FROM posts
WHERE feed_id = $1;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1;

-- name: UpdatePost :exec
UPDATE posts
SET updated_at = $1
WHERE id = $2;

-- name: GetPostsByUser :many
SELECT *
FROM posts
WHERE posts.feed_id IN (
    SELECT feed_follows.feed_id
    FROM feed_follows
    WHERE feed_follows.user_id = $1
)
ORDER BY posts.published_at DESC
LIMIT $2;
 
  
