-- name: GetNewAdditionsStats :one
SELECT * FROM (
  SELECT "id", "created_at", 'event' AS "table", COALESCE("name_en", "name_no") AS "name" FROM "event"
  UNION ALL
  SELECT "id", "created_at", 'category', COALESCE("name_en", "name_no") AS "name" FROM "category"
  UNION ALL
  SELECT "id", "created_at", 'audience', COALESCE("name_en", "name_no") AS "name" FROM "audience"
  UNION ALL
  SELECT "id", "created_at", 'rule', COALESCE("name_en", "name_no") AS "name" FROM "rule"
  UNION ALL
  SELECT "id", "created_at", 'organization', COALESCE("name_en", "name_no") AS "name" FROM "organization"
  UNION ALL
  SELECT "id", "created_at", 'location', COALESCE("name_en", "name_no") AS "name" FROM "location"
  UNION ALL
  SELECT "id", "created_at", 'job_advertisement', COALESCE("title_en", "title_no") AS "name" FROM "job_advertisement"
) AS newest_additions
ORDER BY "created_at" DESC
LIMIT $1;
