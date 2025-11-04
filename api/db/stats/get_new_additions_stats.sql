SELECT * FROM (
  SELECT 
    id, 
    name_en, 
    created_at, 
    updated_at,
    CASE WHEN updated_at > created_at THEN 'updated_at' ELSE 'created_at' END AS order_by,
    'categories' AS source 
  FROM categories
  UNION ALL
  SELECT 
    id, 
    name_no AS name_en, 
    created_at, 
    updated_at,
    CASE WHEN updated_at > created_at THEN 'updated_at' ELSE 'created_at' END AS order_by,
    'events' AS source 
  FROM events
  UNION ALL
  SELECT 
    id, 
    name_en, 
    created_at, 
    updated_at,
    CASE WHEN updated_at > created_at THEN 'updated_at' ELSE 'created_at' END AS order_by,
    'locations' AS source 
  FROM locations
  UNION ALL
  SELECT 
    id, 
    title_en AS name_en, 
    created_at, 
    updated_at,
    CASE WHEN updated_at > created_at THEN 'updated_at' ELSE 'created_at' END AS order_by,
    'jobs' AS source 
  FROM jobs
  UNION ALL
  SELECT 
    id, 
    name_en, 
    created_at, 
    updated_at,
    CASE WHEN updated_at > created_at THEN 'updated_at' ELSE 'created_at' END AS order_by,
    'audiences' AS source 
  FROM audiences
  UNION ALL
  SELECT 
    id, 
    name_en, 
    created_at, 
    updated_at,
    CASE WHEN updated_at > created_at THEN 'updated_at' ELSE 'created_at' END AS order_by,
    'rules' AS source 
  FROM rules
  UNION ALL
  SELECT 
    id, 
    name_en, 
    created_at, 
    updated_at,
    CASE WHEN updated_at > created_at THEN 'updated_at' ELSE 'created_at' END AS order_by,
    'organizations' AS source 
  FROM organizations
) t
ORDER BY GREATEST(created_at, updated_at) DESC
LIMIT 10;