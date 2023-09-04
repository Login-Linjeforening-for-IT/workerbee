-- name: GetRule :one
SELECT * FROM "rule" WHERE "id" = sqlc.arg('id')::int LIMIT 1;
