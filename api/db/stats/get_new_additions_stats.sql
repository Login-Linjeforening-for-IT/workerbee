SELECT * FROM (
  -- categories: always emit a 'created' row; emit an 'updated' row only if updated_at > created_at
  SELECT id,
         name_en,
         created_at,
         created_at AS updated_at,
         'created' AS action,
         'categories' AS source
  FROM categories
  UNION ALL
  SELECT id,
         name_en,
         created_at,
         updated_at,
         'updated' AS action,
         'categories' AS source
  FROM categories
  WHERE updated_at > created_at

  UNION ALL
  SELECT id,
         name_no AS name_en,
         created_at,
         created_at AS updated_at,
         'created' AS action,
         'events' AS source
  FROM events
  UNION ALL
  SELECT id,
         name_no AS name_en,
         created_at,
         updated_at,
         'updated' AS action,
         'events' AS source
  FROM events
  WHERE updated_at > created_at

  UNION ALL
  SELECT id,
         name_en,
         created_at,
         created_at AS updated_at,
         'created' AS action,
         'locations' AS source
  FROM locations
  UNION ALL
  SELECT id,
         name_en,
         created_at,
         updated_at,
         'updated' AS action,
         'locations' AS source
  FROM locations
  WHERE updated_at > created_at

  UNION ALL
  SELECT id,
         title_en AS name_en,
         created_at,
         created_at AS updated_at,
         'created' AS action,
         'jobs' AS source
  FROM jobs
  UNION ALL
  SELECT id,
         title_en AS name_en,
         created_at,
         updated_at,
         'updated' AS action,
         'jobs' AS source
  FROM jobs
  WHERE updated_at > created_at

  UNION ALL
  SELECT id,
         name_en,
         created_at,
         created_at AS updated_at,
         'created' AS action,
         'audiences' AS source
  FROM audiences
  UNION ALL
  SELECT id,
         name_en,
         created_at,
         updated_at,
         'updated' AS action,
         'audiences' AS source
  FROM audiences
  WHERE updated_at > created_at

  UNION ALL
  SELECT id,
         name_en,
         created_at,
         created_at AS updated_at,
         'created' AS action,
         'rules' AS source
  FROM rules
  UNION ALL
  SELECT id,
         name_en,
         created_at,
         updated_at,
         'updated' AS action,
         'rules' AS source
  FROM rules
  WHERE updated_at > created_at

  UNION ALL
  SELECT id,
         name_en,
         created_at,
         created_at AS updated_at,
         'created' AS action,
         'organizations' AS source
  FROM organizations
  UNION ALL
  SELECT id,
         name_en,
         created_at,
         updated_at,
         'updated' AS action,
         'organizations' AS source
  FROM organizations
  WHERE updated_at > created_at
) t
ORDER BY GREATEST(created_at, updated_at) DESC
LIMIT $1;