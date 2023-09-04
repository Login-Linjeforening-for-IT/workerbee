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
    FROM "event" AS e
    INNER JOIN "category" AS c ON e."category" = c."id"
    LEFT OUTER JOIN "location" AS l ON e."location" = l."id"
    WHERE (sqlc.arg(historical)::bool OR ((e."time_end" IS NOT NULL AND e."time_end" > now()) OR (e."time_start" > now() - interval '1 day')))
    ORDER BY e."id"
    LIMIT sqlc.arg('limit')::int
    OFFSET sqlc.arg('offset')::int;


-- name: GetEvent :one
SELECT * FROM "event" WHERE "id" = sqlc.arg('id')::int LIMIT 1;   -- sqlc.embed(event)
