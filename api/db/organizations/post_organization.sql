INSERT INTO organizations
(
    name_no,
    name_en,
    description_no,
    description_en,
    link_homepage,
    link_linkedin,
    link_facebook,
    link_instagram,
    logo
)
VALUES (
    :name_no,
    :name_en,
    :description_no,
    :description_en,
    :link_homepage,
    :link_linkedin,
    :link_facebook,
    :link_instagram,
    :logo
)
RETURNING *;