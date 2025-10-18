INSERT INTO jobs
(
    visible,
    highlight,
    title_no,
    title_en,
    position_title_no,
    position_title_en,
    description_short_no,
    description_short_en,
    description_long_no,
    description_long_en,
    job_type_id,
    time_expire,
    banner_image,
    organization_id,
    application_url
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
)
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
