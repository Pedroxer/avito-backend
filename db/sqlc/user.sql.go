// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: user.sql

package sqlc

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id) values ($1) RETURNING id
`

func (q *Queries) CreateUser(ctx context.Context, id int32) (int32, error) {
	row := q.db.QueryRowContext(ctx, createUser, id)
	err := row.Scan(&id)
	return id, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}
