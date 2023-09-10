-- name: GetAudiencesOfEvent :many
SELECT aud.* FROM "event_audience_relation"
    INNER JOIN "audience" AS aud ON "event_audience_relation"."audience" = aud."id"
    WHERE "event_audience_relation"."event" = sqlc.arg('event_id')::int;

-- name: GetAudiences :many
SELECT "id", "name_no", "name_en", ("deleted_at" IS NOT NULL)::bool AS "is_deleted"
    FROM "audience" ORDER BY "id";

-- name: GetAudience :one
SELECT * FROM "audience" WHERE "id" = sqlc.arg('id')::int LIMIT 1;
