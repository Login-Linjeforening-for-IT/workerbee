-- name: GetJobs :many
SELECT job."id", job."title_no", job."title_en", 
        job."position_title_no", job."position_title_en",
        job."job_type", job."time_publish", job."application_deadline",
        job."application_url", job."updated_at", job."deleted_at", job."visible",
        org."name_no", org."name_en"
    FROM "job_advertisement" AS job
    INNER JOIN "organization" AS org ON job."organization" = org."id"
    ORDER BY job."id"
    LIMIT sqlc.arg('limit')::int
    OFFSET sqlc.arg('offset')::int;

-- name: GetJob :one
SELECT job.*, org."shortname", org."name_no", org."name_en", 
		array(SELECT "skill" FROM "skill" WHERE "ad" = sqlc.arg('id')::int) AS skills,
		array(SELECT "city" FROM "ad_city_relation" WHERE "ad" = sqlc.arg('id')::int) AS cities
    FROM "job_advertisement" AS job
    INNER JOIN "organization" AS org ON job."organization" = org."shortname"
    WHERE job."id" = sqlc.arg('id')::int LIMIT 1;

-- name: CreateJob :one
INSERT INTO "job_advertisement" (
    "visible",
    "title_no", "title_en", 
    "position_title_no", "position_title_en", 
    "description_short_no", "description_short_en", 
    "description_long_no", "description_long_en", 
    "job_type", "time_publish", "application_deadline", "banner_image", 
    "organization", "application_url"
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) RETURNING *;

-- name: UpdateJob :one
UPDATE "job_advertisement"
SET
    "visible" = COALESCE(sqlc.narg(visible), visible),
    "title_no" = COALESCE(sqlc.narg(title_no), title_no),
    "title_en" = COALESCE(sqlc.narg(title_en), title_en),
    "position_title_no" = COALESCE(sqlc.narg(position_title_no), position_title_no),
    "position_title_en" = COALESCE(sqlc.narg(position_title_en), position_title_en),
    "description_short_no" = COALESCE(sqlc.narg(description_short_no), description_short_no),
    "description_short_en" = COALESCE(sqlc.narg(description_short_en), description_short_en),
    "description_long_no" = COALESCE(sqlc.narg(description_long_no), description_long_no),
    "description_long_en" = COALESCE(sqlc.narg(description_long_en), description_long_en),
    "job_type" = COALESCE(sqlc.narg(job_type), job_type),
    "time_publish" = COALESCE(sqlc.narg(time_publish), time_publish),
    "application_deadline" = COALESCE(sqlc.narg(application_deadline), application_deadline),
    "banner_image" = COALESCE(sqlc.narg(banner_image), banner_image),
    "organization" = COALESCE(sqlc.narg(organization), organization),
    "application_url" = COALESCE(sqlc.narg(application_url), application_url),
    "updated_at" = now()
WHERE "id" = sqlc.arg('id')::int RETURNING *;
