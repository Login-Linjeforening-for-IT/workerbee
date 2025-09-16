-- name: post_event :one
INSERT INTO events
(
    visible,
    name_no,
    name_en,
    description_no,
    description_en,
    informational_no,
    informational_en,
    time_type,
    time_start,
    time_end,
    time_publish,
    time_signup_release,
    time_signup_deadline,
    canceled,
    digital,
    highlight,
    image_small,
    image_banner,
    link_facebook,
    link_discord,
    link_signup,
    link_stream,
    capacity,
    full,
    category_id,
    organization_id,
    location_id,
    parent_id,
    rule_id,
    audience_id
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30)
RETURNING *;