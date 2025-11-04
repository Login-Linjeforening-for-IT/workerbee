SELECT * FROM (
  SELECT 
    id, 
    name_en, 
    created_at, 
    updated_at,
    CASE WHEN updated_at > created_at THEN 'updated' ELSE 'created' END AS action,
    'categories' AS source 
  FROM categories
  UNION ALL
  SELECT 
    id, 
    name_no AS name_en, 
    created_at, 
    updated_at,
    CASE WHEN updated_at > created_at THEN 'updated' ELSE 'created' END AS action,
    'events' AS source 
  FROM events
  UNION ALL
  SELECT 
    id, 
    name_en, 
    created_at, 
    updated_at,
    CASE WHEN updated_at > created_at THEN 'updated' ELSE 'created' END AS action,
    'locations' AS source 
  FROM locations
  UNION ALL
  SELECT 
    id, 
    title_en AS name_en, 
    created_at, 
    updated_at,
    CASE WHEN updated_at > created_at THEN 'updated' ELSE 'created' END AS action,
    'jobs' AS source 
  FROM jobs
  UNION ALL
  SELECT 
    id, 
    name_en, 
    created_at, 
    updated_at,
    CASE WHEN updated_at > created_at THEN 'updated' ELSE 'created' END AS action,
    'audiences' AS source
  FROM audiences
  UNION ALL
  SELECT 
    id, 
    name_en, 
    created_at, 
    updated_at,
    CASE WHEN updated_at > created_at THEN 'updated' ELSE 'created' END AS action,
    'rules' AS source 
  FROM rules
  UNION ALL
  SELECT 
    id, 
    name_en, 
    created_at, 
    updated_at,
    CASE WHEN updated_at > created_at THEN 'updated' ELSE 'created' END AS action,
    'organizations' AS source 
  FROM organizations
) t
ORDER BY GREATEST(created_at, updated_at) DESC
LIMIT 10;