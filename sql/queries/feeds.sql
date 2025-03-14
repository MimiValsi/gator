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
--

-- name: GetFeeds :many
SELECT id, created_at, updated_at, name, url, name, user_id, last_fetched_at
  FROM feeds;
--

-- name: GetFeedByURL :one
SELECT id, created_at, updated_at, name, url, user_id, last_fetched_at
  FROM feeds
WHERE url = $1;
--

-- name: MarkFeedFetched :exec
UPDATE feeds
  SET updated_at = NOW(),
      last_fetched_at = NOW()
WHERE id = $1
RETURNING *;
--

-- name: GetNextFeedtoFetch :one
SELECT *
  FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;
--
