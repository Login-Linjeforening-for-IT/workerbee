-- name: get_job :one
-- name: get_jobs :many
SELECT
	ja.*,
	array_agg(DISTINCT c.name) AS cities,
	array_agg(DISTINCT s.name) AS skills
FROM job_advertisements ja
LEFT JOIN ad_city_relation acr ON ja.id = acr.job_advertisement_id
LEFT JOIN cities c ON acr.city_id = c.id
LEFT JOIN ad_skill_relation asr ON ja.id = asr.job_advertisement_id
LEFT JOIN skills s ON asr.skill_id = s.id
WHERE ja.id = $1
GROUP BY ja.id;