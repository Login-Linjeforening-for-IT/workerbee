SELECT
	e.id,
	e.name_no,
	e.name_en,
	e.description_no,
	e.description_en,
	e.time_start,
	e.time_end
FROM events AS e
WHERE
    e.visible = TRUE
    AND e.time_end > NOW()
    AND (
        cardinality($1::int[]) = 0
        OR e.category_id = ANY($1::int[])
    );
