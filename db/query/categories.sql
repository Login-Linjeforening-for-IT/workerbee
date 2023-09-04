-- name: GetCategory :one
SELECT * FROM "category" WHERE "id" = sqlc.arg('id')::int LIMIT 1;
