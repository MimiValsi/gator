-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
  $1,
  $2,
  $3,
  $4
)
RETURNING *;

-- name: GetUser :one
SELECT id, created_at, updated_at, name
  FROM users
WHERE name = $1;

-- name: GetUserByID :one
SELECT * FROM users where id = $1;

-- name: DeleteUsers :exec
DELETE FROM users CASCADE;

-- name: GetUsers :many
SELECT name
FROM users;
