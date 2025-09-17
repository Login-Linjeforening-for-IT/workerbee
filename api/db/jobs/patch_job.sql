-- name: patch_job :one
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
RETURNING *;