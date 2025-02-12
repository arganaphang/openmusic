-- name: GetSongs :many
SELECT * FROM "songs";

-- name: GetSongByID :one
SELECT * FROM "songs" WHERE "id" = $1;

-- name: GetSongByAlbumID :many
SELECT * FROM "songs" WHERE "album_id" = $1;

-- name: CreateSong :one
INSERT INTO "songs" ("id", "title", "year", "genre", "performer", "duration", "album_id") VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- name: UpdateSong :one
UPDATE "songs" SET "title" = $1, "year" = $2, "genre" = $3, "performer" = $4, "duration" = $5, "album_id" = $6 WHERE "id" = $7 RETURNING *;

-- name: DeleteSong :exec
DELETE FROM "songs" WHERE "id" = $1;
