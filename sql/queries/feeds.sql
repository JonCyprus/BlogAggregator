-- name: CreateFeed :one
INSERT INTO feeds (created_at, updated_at, name, url, user_id)
VALUES (
           $1,
           $2,
           $3,
           $4,
         $5
       )
RETURNING *;

-- name: GetFeed :one
SELECT * FROM feeds
WHERE name = $1;

-- name: GetFeedByURL :one
SELECT * FROM feeds
WHERE url = $1;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeedsForPrint :many
SELECT feeds.name, feeds.url, users.name AS username
FROM feeds
INNER JOIN users ON feeds.user_id = users.id;

-- name: ResetFeedsTable :exec
DELETE FROM feeds;

-- name: MarkFeedFetched :one
UPDATE feeds
SET updated_at = NOW(),
    last_fetched_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetNextFeedToFetch :one
SELECT id, url FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;