-- name: GetAudiencesOfEvent :many
SELECT aud.* FROM "event_audience_relation"
    INNER JOIN "audience" AS aud ON "event_audience_relation"."audience" = aud."id"
    WHERE "event_audience_relation"."event" = sqlc.arg('event_id')::int;
