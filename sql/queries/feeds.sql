-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeedByID :one
SELECT * FROM feeds WHERE id = $1;

-- name: CreateFeed :one
INSERT INTO feeds (id, url, name, user_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateFeed :one
UPDATE feeds
SET
    url = $1,
    name = $2,
    updated_at = NOW()
WHERE id = $3
RETURNING *;

-- name: DeleteFeed :exec
DELETE FROM feeds WHERE id = $1;
