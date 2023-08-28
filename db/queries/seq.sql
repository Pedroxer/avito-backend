-- name: CreateSeg :one
INSERT INTO "segments" (name) values ($1) RETURNING *;

-- name: UpdateSeg :exec
UPDATE "segments" SET name = $2 where id = $1;

-- name: DeleteSeg :exec
DELETE FROM "segments" where id = $1;

-- name: GetSeq :one
SELECT * FROm "segments" where id = $1;