DELETE FROM categories
WHERE id = $1
RETURNING id;