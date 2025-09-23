-- name: get_jobs :many
SELECT
	ja.*,
	array_agg(DISTINCT c.name) FILTER (WHERE c.name IS NOT NULL) AS cities,
	array_agg(DISTINCT s.name) FILTER (WHERE s.name IS NOT NULL) AS skills,
	COUNT(*) OVER() AS total_count
FROM job_advertisements ja
LEFT JOIN ad_city_relation acr ON ja.id = acr.job_advertisement_id
LEFT JOIN cities c ON acr.city_id = c.id
LEFT JOIN ad_skill_relation asr ON ja.id = asr.job_advertisement_id
LEFT JOIN skills s ON asr.skill_id = s.id
WHERE (
	$1 = '' OR
	to_json(ja)::text ILIKE '%' || $1 || '%'
)
GROUP BY ja.id
LIMIT $2 OFFSET $3;