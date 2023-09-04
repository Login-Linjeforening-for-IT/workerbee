-- name: GetLocation :one
SELECT * FROM "location" WHERE "id" = sqlc.arg('id')::int LIMIT 1;
