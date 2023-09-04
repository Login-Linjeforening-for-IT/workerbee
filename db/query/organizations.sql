-- name: GetOrganizationsOfEvent :many
SELECT org.* FROM "event_organization_relation"
    INNER JOIN "organization" AS org ON "event_organization_relation"."organization" = org."shortname"
    WHERE "event_organization_relation"."event" = sqlc.arg('event_id')::int;
