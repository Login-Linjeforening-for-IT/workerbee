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
    job_type,
    time_expire,
    application_deadline,
    banner_image,
    organization_id,
    application_url
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11,
    $12,
    $13,
    $14,
    $15,
    $16
)
RETURNING id;