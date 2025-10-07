-- name: post_organization :one
INSERT INTO organizations
(
    shortname,
    name_no,
    name_en,
    description_no,
    description_en,
    type,
    link_homepage,
    link_linkedin,
    link_facebook,
    link_instagram,
    logo
)
VALUES (
    :shortname,
    :name_no,
    :name_en,
    :description_no,
    :description_en,
    :type,
    :link_homepage,
    :link_linkedin,
    :link_facebook,
    :link_instagram,
    :logo
)
RETURNING *;