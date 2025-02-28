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
select id, created_at, updated_at, name
from users
where name = $1
;

-- name: TruncateUsers :exec
TRUNCATE TABLE users CASCADE;

-- name: GetUsers :many
select name
from users
;

