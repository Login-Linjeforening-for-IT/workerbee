-- name: put_job :one
UPDATE jobs
SET
    visible = $2,
    highlight = $3,
    title_no = $4,
    title_en = $5,
    position_title_no = $6,
    position_title_en = $7,
    description_short_no = $8,
    description_short_en = $9,
    description_long_no = $10,
    description_long_en = $11,
    job_type = $12,
    time_publish = $13,
    time_expire = $14,
    application_deadline = $15,
    banner_image = $16,
    organization_id = $17,
    application_url = $18,
    updated_at = NOW()
WHERE id = $1
RETURNING *,
(
    SELECT array_agg(DISTINCT c.name) 
    FROM ad_city_relation acr
    JOIN cities c ON acr.city_id = c.id
    WHERE acr.job_id = jobs.id
) AS cities,
(
    SELECT array_agg(DISTINCT s.name) 
    FROM ad_skill_relation asr
    JOIN skills s ON asr.skill_id = s.id
    WHERE asr.job_id = jobs.id
) AS skills;