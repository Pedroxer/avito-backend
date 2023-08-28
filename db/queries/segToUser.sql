-- name: GetSegOfUser :many
SELECT * FROM "seg_to_user" WHERE user_id = $1;

-- name: AddSegToUser :one
INSERT INTO "seg_to_user" (user_id, seg_id) VALUES ($1, $2) RETURNING *;

-- name: RemoveSegFromUser :exec
DELETE FROM "seg_to_user" WHERE user_id = $1 and seg_id = $2;

