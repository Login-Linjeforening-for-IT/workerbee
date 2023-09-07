-- name: GetEvents :many
SELECT e."id", e."visible",
        e."name_no", e."name_en",
        e."time_type", e."time_start", e."time_end", e."time_publish",
        e."canceled", e."link_signup", e."capacity", e."full",
        c."name_no" AS category_name_no, c."name_en"  AS category_name_en, 
        -- TODO: Add audience
        l."name_no" AS location_name_no, l."name_en" AS location_name_en,
        -- TODO: Add organizer
        e."updated_at", e."deleted_at" IS NOT NULL AS is_deleted
        -- deleted_at == null => is_deleted == false
    FROM "event" AS e
    INNER JOIN "category" AS c ON e."category" = c."id"
    LEFT OUTER JOIN "location" AS l ON e."location" = l."id"
    WHERE (sqlc.arg(historical)::bool OR ((e."time_end" IS NOT NULL AND e."time_end" > now()) OR (e."time_start" > now() - interval '1 day')))
    ORDER BY e."id"
    LIMIT sqlc.arg('limit')::int
    OFFSET sqlc.arg('offset')::int;


-- name: GetEvent :one
SELECT * FROM "event" WHERE "id" = sqlc.arg('id')::int LIMIT 1;   -- sqlc.embed(event)

-- name: CreateEvent :one
INSERT INTO "event" (
    "visible",
    "name_no", "name_en",
    "description_no", "description_en",
    "informational_no", "informational_en",
    "time_type", "time_start", "time_end", "time_publish",
    "time_signup_release", "time_signup_deadline",
    "canceled", "digital", "highlight",
    "image_small", "image_banner",
    "link_facebook", "link_discord", "link_signup", "link_stream",
    "capacity", "full",
    "category", "location", "parent", "rule"
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28)
RETURNING *;

-- name: UpdateEvent :one
UPDATE "event"
SET
    "visible" = COALESCE(sqlc.narg(visible), visible),
    "name_no" = COALESCE(sqlc.narg(name_no), name_no),
    "name_en" = COALESCE(sqlc.narg(name_en), name_en),
    "description_no" = COALESCE(sqlc.narg(description_no), description_no),
    "description_en" = COALESCE(sqlc.narg(description_en), description_en),
    "informational_no" = COALESCE(sqlc.narg(informational_no), informational_no),
    "informational_en" = COALESCE(sqlc.narg(informational_en), informational_en),
    "time_type" = COALESCE(sqlc.narg(time_type), time_type),
    "time_start" = COALESCE(sqlc.narg(time_start), time_start),
    "time_end" = COALESCE(sqlc.narg(time_end), time_end),
    "time_publish" = COALESCE(sqlc.narg(time_publish), time_publish),
    "time_signup_release" = COALESCE(sqlc.narg(time_signup_release), time_signup_release),
    "time_signup_deadline" = COALESCE(sqlc.narg(time_signup_deadline), time_signup_deadline),
    "canceled" = COALESCE(sqlc.narg(canceled), canceled),
    "digital" = COALESCE(sqlc.narg(digital), digital),
    "highlight" = COALESCE(sqlc.narg(highlight), highlight),
    "image_small" = COALESCE(sqlc.narg(image_small), image_small),
    "image_banner" = COALESCE(sqlc.narg(image_banner), image_banner),
    "link_facebook" = COALESCE(sqlc.narg(link_facebook), link_facebook),
    "link_discord" = COALESCE(sqlc.narg(link_discord), link_discord),
    "link_signup" = COALESCE(sqlc.narg(link_signup), link_signup),
    "link_stream" = COALESCE(sqlc.narg(link_stream), link_stream),
    "capacity" = COALESCE(sqlc.narg(capacity), capacity),
    "full" = COALESCE(sqlc.narg('full'), "full"),
    "category" = COALESCE(sqlc.narg(category), category),
    "location" = COALESCE(sqlc.narg('location'), "location"),
    "parent" = COALESCE(sqlc.narg(parent), parent),
    "rule" = COALESCE(sqlc.narg(rule), rule),
    "updated_at" = now()
WHERE "id" = sqlc.arg(id)::int
RETURNING *;

-- name: SoftDeleteEvent :one
UPDATE "event"
SET
    "deleted_at" = now()
WHERE "id" = sqlc.arg(id)::int
RETURNING *;

-- name: AddOrganizationToEvent :exec
INSERT INTO "event_organization_relation" ("event", "organization") VALUES ($1, $2) ON CONFLICT DO NOTHING;

-- name: AddAudienceToEvent :exec
INSERT INTO "event_audience_relation" ("event", "audience") VALUES ($1, $2) ON CONFLICT DO NOTHING;
