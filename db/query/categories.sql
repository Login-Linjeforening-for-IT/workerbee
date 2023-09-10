-- name: GetCategory :one
SELECT * FROM "category" WHERE "id" = sqlc.arg('id')::int LIMIT 1;

-- name: GetCategories :many
SELECT "id", "color", "name_no", "name_en", ("deleted_at" IS NOT NULL)::bool AS "is_deleted"
    FROM "category" ORDER BY "id";
