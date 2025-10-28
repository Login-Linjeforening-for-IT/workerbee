UPDATE honey
SET text = $1
WHERE service = $2
    AND page = $3
    AND language = $4;