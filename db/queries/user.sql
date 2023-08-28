-- name: CreateUser :one
INSERT INTO users (name) values ($1) RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE name = $1;