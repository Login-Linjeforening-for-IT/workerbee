UPDATE organizations
SET
    name_no = :name_no,
    name_en = :name_en,
    description_no = :description_no,
    description_en = :description_en,
    link_homepage = :link_homepage,
    link_linkedin = :link_linkedin,
    link_facebook = :link_facebook,
    link_instagram = :link_instagram,
    logo = :logo,
    updated_at = NOW()
WHERE id = :id
RETURNING *;
