-- name: GetAlbums :many
SELECT * FROM "albums";

-- name: GetAlbumByID :one
SELECT * FROM "albums" WHERE "id" = $1;

-- name: CreateAlbum :one
INSERT INTO "albums" ("id", "name", "year") VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateAlbum :one
UPDATE "albums" SET "name" = $1, "year" = $2 WHERE "id" = $3 RETURNING *;

-- name: DeleteAlbum :exec
DELETE FROM "albums" WHERE "id" = $1;
