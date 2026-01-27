INSERT 
INTO quotes (author, quoted, content)
VALUES (:author, :quoted, :content)
RETURNING *;