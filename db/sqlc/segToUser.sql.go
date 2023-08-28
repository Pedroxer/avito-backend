// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: segToUser.sql

package sqlc

import (
	"context"
	"database/sql"
)

const addSegToUser = `-- name: AddSegToUser :one
INSERT INTO "seg_to_user" (user_id, seg_id) VALUES ($1, $2) RETURNING user_id, seg_id
`

type AddSegToUserParams struct {
	UserID sql.NullInt32 `json:"user_id"`
	SegID  sql.NullInt32 `json:"seg_id"`
}

func (q *Queries) AddSegToUser(ctx context.Context, arg AddSegToUserParams) (SegToUser, error) {
	row := q.db.QueryRowContext(ctx, addSegToUser, arg.UserID, arg.SegID)
	var i SegToUser
	err := row.Scan(&i.UserID, &i.SegID)
	return i, err
}

const getSegOfUser = `-- name: GetSegOfUser :many
SELECT user_id, seg_id FROM "seg_to_user" WHERE user_id = $1
`

func (q *Queries) GetSegOfUser(ctx context.Context, userID sql.NullInt32) ([]SegToUser, error) {
	rows, err := q.db.QueryContext(ctx, getSegOfUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SegToUser
	for rows.Next() {
		var i SegToUser
		if err := rows.Scan(&i.UserID, &i.SegID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeSegFromUser = `-- name: RemoveSegFromUser :exec
DELETE FROM "seg_to_user" WHERE user_id = $1 and seg_id = $2
`

type RemoveSegFromUserParams struct {
	UserID sql.NullInt32 `json:"user_id"`
	SegID  sql.NullInt32 `json:"seg_id"`
}

func (q *Queries) RemoveSegFromUser(ctx context.Context, arg RemoveSegFromUserParams) error {
	_, err := q.db.ExecContext(ctx, removeSegFromUser, arg.UserID, arg.SegID)
	return err
}