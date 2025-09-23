-- name: GetNewAdditionsStats :one
SELECT * FROM (
  SELECT "id", "created_at", 'events' AS "table", COALESCE("name_en", "name_no") AS "name" FROM "events"
  UNION ALL
  SELECT "id", "created_at", 'categories', COALESCE("name_en", "name_no") AS "name" FROM "categories"
  UNION ALL
  SELECT "id", "created_at", 'audiences', COALESCE("name_en", "name_no") AS "name" FROM "audiences"
  UNION ALL
  SELECT "id", "created_at", 'rules', COALESCE("name_en", "name_no") AS "name" FROM "rules"
  UNION ALL
  SELECT "id", "created_at", 'organizations', COALESCE("name_en", "name_no") AS "name" FROM "organizations"
  UNION ALL
  SELECT "id", "created_at", 'locations', COALESCE("name_en", "name_no") AS "name" FROM "locations"
  UNION ALL
  SELECT "id", "created_at", 'jobs', COALESCE("title_en", "title_no") AS "name" FROM "jobs"
) AS newest_additions
ORDER BY "created_at" DESC
LIMIT $1;
