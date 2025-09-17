-- name: patch_organization :one
UPDATE organizations
SET
    shortname = $2,
    name_no = $3,
    name_en = $4,
    description_no = $5,
    description_en = $6,
    type = $7,
    link_homepage = $8,
    link_linkedin = $9,
    link_facebook = $10,
    link_instagram = $11,
    logo = $12,
    updated_at = NOW()
WHERE id = $1
RETURNING *;
