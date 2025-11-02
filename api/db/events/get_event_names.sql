-- name: GetEventNames :many
SELECT
    id,
    name_en,
    time_start
FROM events
ORDER BY time_start DESC;