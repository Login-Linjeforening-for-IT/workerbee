-- name: GetRule :one
SELECT * FROM "rule" WHERE "id" = sqlc.arg('id')::int LIMIT 1;

-- name: GetRules :many
SELECT "id", "name_no", "name_en", "updated_at" FROM "rule"
    LIMIT sqlc.arg('limit')::int
    OFFSET sqlc.arg('offset')::int;

-- name: CreateRule :one
INSERT INTO "rule" (
    "name_no", "name_en",
    "description_no", "description_en"
) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdateRule :one
UPDATE "rule"
SET
    "name_no" = COALESCE(sqlc.narg(name_no), name_no),
    "name_en" = COALESCE(sqlc.narg(name_en), name_en),
    "description_no" = COALESCE(sqlc.narg(description_no), description_no),
    "description_en" = COALESCE(sqlc.narg(description_en), description_en),
    "updated_at" = now()
WHERE "id" = sqlc.arg(id)::int
RETURNING *;

-- name: SoftDeleteRule :one
UPDATE "rule"
SET
    "deleted_at" = now()
WHERE "id" = sqlc.arg('id')::int
RETURNING *;
