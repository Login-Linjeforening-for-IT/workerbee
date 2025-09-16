-- name: get_rule :one
SELECT * FROM rules WHERE id = $1;