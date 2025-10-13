-- name: GetEvents :many
SELECT
    e.id,
    e.name_en,
    e.name_no,
    e.time_start,
    e.time_end,
    c.id AS "category.id",
    c.name_no AS "category.name_no",
    c.name_en AS "category.name_en",
    l.id AS "location.id",
    l.name_no AS "location.name_no",
    l.name_en AS "location.name_en",
    o.id AS "organization.id",
    o.name_no AS "organization.name_no",
    o.name_en AS "organization.name_en"
FROM events AS e
LEFT JOIN categories AS c ON e.category_id = c.id
LEFT JOIN locations AS l ON e.location_id = l.id
LEFT JOIN organizations AS o ON e.organization_id = o.id
WHERE (
        $2::bool
        OR (
            (
                e.time_end IS NOT NULL
                AND e.time_end > now()
            )
            OR (e.time_start > now() - interval '1 day')
        )
    )
    AND (
        $1 = ''
        OR to_json(e)::text ILIKE '%' || $1 || '%'
    )
    AND (
        cardinality($3::int[]) = 0
        OR e.category_id = ANY($3)
    )