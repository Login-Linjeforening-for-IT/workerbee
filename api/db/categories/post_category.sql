INSERT INTO categories (
    name_no, 
    name_en, 
    color
)
VALUES (
    :name_no, 
    :name_en, 
    :color
)
RETURNING *;