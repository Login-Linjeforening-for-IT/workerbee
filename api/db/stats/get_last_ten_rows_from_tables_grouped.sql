SELECT
  json_build_object(
    'categories', (
      SELECT json_agg(row_to_json(t))
      FROM (
        SELECT id, name_en, created_at
        FROM categories
        ORDER BY created_at DESC
        LIMIT 10
      ) t
    ),
    'events', (
      SELECT json_agg(row_to_json(t))
      FROM (
        SELECT id, name_no AS name_en, created_at
        FROM events
        ORDER BY created_at DESC
        LIMIT 10
      ) t
    ),
    'locations', (
      SELECT json_agg(row_to_json(t))
      FROM (
        SELECT id, name_en, created_at
        FROM locations
        ORDER BY created_at DESC
        LIMIT 10
      ) t
    ),
    'jobs', (
      SELECT json_agg(row_to_json(t))
      FROM (
        SELECT id, title_en AS name_en, created_at
        FROM jobs
        ORDER BY created_at DESC
        LIMIT 10
      ) t
    ),
    'audiences', (
      SELECT json_agg(row_to_json(t))
      FROM (
        SELECT id, name_en, created_at
        FROM audiences
        ORDER BY created_at DESC
        LIMIT 10
      ) t
    ),
    'rules', (
      SELECT json_agg(row_to_json(t))
      FROM (
        SELECT id, name_en, created_at
        FROM rules
        ORDER BY created_at DESC
        LIMIT 10
      ) t
    ),
    'organizations', (
      SELECT json_agg(row_to_json(t))
      FROM (
        SELECT id, name_en, created_at
        FROM organizations
        ORDER BY created_at DESC
        LIMIT 10
      ) t
    )
  ) AS grouped_data;
