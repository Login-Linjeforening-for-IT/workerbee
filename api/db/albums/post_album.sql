INSERT INTO albums 
( 
    name_en, 
    name_no,
    description_en,
    description_no,
    year,
    event_id
) VALUES 
(
    :name_en,
    :name_no,
    :description_en,
    :description_no,
    :year,
    :event_id
)
RETURNING *;