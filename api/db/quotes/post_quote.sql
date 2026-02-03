INSERT 
INTO quotes (author, quoted, content)
VALUES (:author, :quoted, :content)
RETURNING id, quoted, content, created_at, updated_at;