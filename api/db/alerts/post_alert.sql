INSERT INTO alerts 
(   
    service, 
    page, 
    title_en, 
    title_no, 
    description_en, 
    description_no
)
VALUES (
    :service, 
    :page, 
    :title_en, 
    :title_no, 
    :description_en, 
    :description_no
)
RETURNING *;