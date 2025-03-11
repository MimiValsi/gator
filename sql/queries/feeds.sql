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
SELECT f.name, f.url, u.name
  FROM feeds AS f
LEFT JOIN users AS u 
  ON f.user_id = u.id;

-- name: GetFeed :one
SELECT id, created_at, updated_at, name, url, user_id
  FROM feeds
WHERE url = $1;

-- name: CreateFeedFollow :many
WITH inserted_feed_follow AS (
  INSERT INTO feed_follows (
    id,
    created_at,
    updated_at,
    user_id,
    feed_id
  )
  VALUES ($1, $2, $3, $4, $5)
  RETURNING *
)
SELECT
  inserted_feed_follow.*,
  feeds.name AS feed_name,
  users.name AS user_name
FROM inserted_feed_follow
INNER JOIN feeds ON feeds.id = inserted_feed_follow.feed_id
INNER JOIN users ON users.id = inserted_feed_follow.user_id;

-- name: GetFollowingUsers :many
SELECT feeds.name AS feed_name
  FROM feed_follows
INNER JOIN feeds ON feeds.id = feed_follows.feed_id
  WHERE feed_follows.user_id = $1;

-- name: DeleteFeedFollows :exec
DELETE FROM feed_follows CASCADE;
