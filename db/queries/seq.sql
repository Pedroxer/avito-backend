-- name: CreateSeg :one
INSERT INTO "segments" (name) values ($1) RETURNING *;

-- name: UpdateSeg :exec
UPDATE "segments" SET name = $2 where id = $1;

-- name: DeleteSeg :exec
DELETE FROM "segments" where id = $1;

-- name: GetSeg :one
SELECT * FROM "segments" where id = $1;