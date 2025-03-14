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
SELECT f.name, f.url, u.name
  FROM feeds AS f
LEFT JOIN users AS u 
  ON f.user_id = u.id;
--

-- name: GetFeedByURL :one
SELECT id, created_at, updated_at, name, url, user_id
  FROM feeds
WHERE url = $1;
--

-- name: MarkFeedFetched :exec
UPDATE feeds
  SET updated_at = $1,
      last_fetched_at = $2
WHERE id = $3;
--

-- name: GetNextFeedtoFetch :one
SELECT id, created_at, updated_at, name, url, user_id, last_fetched_at
  FROM feeds
ORDER BY updated_at ASC NULLS FIRST;
--
