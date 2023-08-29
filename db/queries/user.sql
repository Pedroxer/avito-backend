-- name: CreateUser :one
INSERT INTO users (id) values ($1) RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;