-- name: GetOrganizationsOfEvent :many
SELECT org.*, ("deleted_at" IS NOT NULL)::bool AS "is_deleted" FROM "event_organization_relation"
    INNER JOIN "organization" AS org ON "event_organization_relation"."organization" = org."shortname"
    WHERE "event_organization_relation"."event" = sqlc.arg('event_id')::int;

-- name: GetOrganizations :many
SELECT "shortname", "name_no", "name_en", "link_homepage", "logo", "updated_at", ("deleted_at" IS NOT NULL)::bool AS "is_deleted"
FROM "organization"
ORDER BY "shortname"
LIMIT sqlc.arg('limit')::int
OFFSET sqlc.arg('offset')::int;

-- name: GetOrganization :one
SELECT * FROM "organization" WHERE "shortname" = sqlc.arg('shortname')::text LIMIT 1;

-- name: CreateOrganization :one
INSERT INTO "organization" (
    "shortname",
    "name_no", "name_en",
    "description_no", "description_en",
    "type",
    "link_homepage", "link_linkedin", "link_facebook", "link_instagram",
    "logo"
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING *;

-- name: UpdateOrganization :one
UPDATE "organization"
SET
    "name_no" = COALESCE(sqlc.narg(name_no), name_no),
    "name_en" = COALESCE(sqlc.narg(name_en), name_en),
    "description_no" = COALESCE(sqlc.narg(description_no), description_no),
    "description_en" = COALESCE(sqlc.narg(description_en), description_en),
    "type" = COALESCE(sqlc.narg('type'), "type"),
    "link_homepage" = COALESCE(sqlc.narg(link_homepage), link_homepage),
    "link_linkedin" = COALESCE(sqlc.narg(link_linkedin), link_linkedin),
    "link_facebook" = COALESCE(sqlc.narg(link_facebook), link_facebook),
    "link_instagram" = COALESCE(sqlc.narg(link_instagram), link_instagram),
    "logo" = COALESCE(sqlc.narg(logo), logo),
    "updated_at" = now()
WHERE "shortname" = sqlc.arg('shortname')::text RETURNING *;

-- name: SoftDeleteOrganization :one
UPDATE "organization"
SET
    "deleted_at" = now(),
    "updated_at" = now()
WHERE "shortname" = sqlc.arg('shortname')::text RETURNING *;
