UPDATE quotes
SET 
    author = :author, 
    quoted = :quoted, 
    content = :content, 
    updated_at = NOW()
WHERE id = :id
RETURNING id, quoted, content, created_at, updated_at;