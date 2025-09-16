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
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING *;