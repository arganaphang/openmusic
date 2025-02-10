-- name: GetUsers :many
SELECT * FROM "users";

-- name: GetUserByID :one
SELECT * FROM "users" WHERE "id" = $1;

-- name: CreateUser :one
INSERT INTO "users" ("name", "email", "password") VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateUser :one
UPDATE "users" SET "name" = $1, "email" = $2, "password" = $3 WHERE "id" = $4 RETURNING *;

-- name: DeleteUser :exec
DELETE FROM "users" WHERE "id" = $1;
