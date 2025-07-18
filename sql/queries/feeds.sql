-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT *
FROM feeds;

-- name: GetFeedByURL :one
SELECT *
FROM feeds
WHERE url = $1
LIMIT 1;

-- name: DeleteALLFeeds :exec
DELETE FROM feeds;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetch_at = $1, updated_at = $1
WHERE id = $2; 

-- name: GetNextFeedToFetch :one
SELECT *
FROM feeds
INNER JOIN feed_follows ON feeds.id = feed_follows.feed_id AND feed_follows.user_id = $1
ORDER BY last_fetch_at  ASC NULLS FIRST
LIMIT 1;

-- 
