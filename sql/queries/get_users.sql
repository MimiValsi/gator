-- name: GetUser :one
SELECT id, created_at, updated,at, name FROM users WHERE name = $1;
