SELECT * FROM (
  SELECT id, name_en, created_at, 'categories' AS source FROM categories
  UNION ALL
  SELECT id, name_no AS name_en, created_at, 'events' AS source FROM events
  UNION ALL
  SELECT id, name_en, created_at, 'locations' AS source FROM locations
  UNION ALL
  SELECT id, title_en AS name_en, created_at, 'jobs' AS source FROM jobs
  UNION ALL
  SELECT id, name_en, created_at, 'audiences' AS source FROM audiences
  UNION ALL
  SELECT id, name_en, created_at, 'rules' AS source FROM rules
  UNION ALL
  SELECT id, name_en, created_at, 'organizations' AS source FROM organizations
) t
ORDER BY created_at DESC
LIMIT 10;