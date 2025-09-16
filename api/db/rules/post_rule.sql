-- name: post_rule :one
INSERT INTO rules
(
    name_no,
    name_en,
    description_no,
    description_en
)
VALUES ($1, $2, $3, $4)
RETURNING *;