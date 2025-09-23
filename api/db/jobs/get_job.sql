-- name: get_job :one
SELECT
    ja.*,
    city_agg.cities,
    skill_agg.skills
FROM job_advertisements ja
LEFT JOIN (
    SELECT acr.job_advertisement_id,
           array_agg(DISTINCT c.name) FILTER (WHERE c.name IS NOT NULL) AS cities
    FROM ad_city_relation acr
    LEFT JOIN cities c ON acr.city_id = c.id
    GROUP BY acr.job_advertisement_id
) city_agg ON city_agg.job_advertisement_id = ja.id
LEFT JOIN (
    SELECT asr.job_advertisement_id,
           array_agg(DISTINCT s.name) FILTER (WHERE s.name IS NOT NULL) AS skills
    FROM ad_skill_relation asr
    LEFT JOIN skills s ON asr.skill_id = s.id
    GROUP BY asr.job_advertisement_id
) skill_agg ON skill_agg.job_advertisement_id = ja.id
WHERE ja.id = $1;