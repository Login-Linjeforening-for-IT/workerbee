-- name: get_job :one
SELECT
    ja.id,
    ja.visible,
    ja.highlight,
    ja.title_en,
    ja.title_no,
    ja.description_short_en,
    ja.description_short_no,
    ja.description_long_en,
    ja.description_long_no,
    ja.position_title_en,
    ja.position_title_no,
    ja.job_type,
    ja.time_publish,
    ja.time_expire,
    ja.banner_image,
    ja.application_url,
    ja.created_at,
    ja.updated_at,
    org.id AS "organization.id",
    org.name_no AS "organization.name_no",
    org.name_en AS "organization.name_en",
    org.description_no AS "organization.description_no",
    org.description_en AS "organization.description_en",
    org.link_homepage AS "organization.link_homepage",
    org.link_facebook AS "organization.link_facebook",
    org.link_linkedin AS "organization.link_linkedin",
    org.created_at AS "organization.created_at",
    org.updated_at AS "organization.updated_at",
    org.logo AS "organization.logo",
    city_agg.cities,
    skill_agg.skills
FROM jobs ja
JOIN organizations org ON ja.organization_id = org.id
LEFT JOIN (
    SELECT acr.job_id,
           array_agg(DISTINCT c.name) FILTER (WHERE c.name IS NOT NULL) AS cities
    FROM ad_city_relation acr
    LEFT JOIN cities c ON acr.city_id = c.id
    GROUP BY acr.job_id
) city_agg ON city_agg.job_id = ja.id
LEFT JOIN (
    SELECT asr.job_id,
           array_agg(DISTINCT s.name) FILTER (WHERE s.name IS NOT NULL) AS skills
    FROM ad_skill_relation asr
    LEFT JOIN skills s ON asr.skill_id = s.id
    GROUP BY asr.job_id
) skill_agg ON skill_agg.job_id = ja.id
WHERE ja.id = $1
    AND ja.time_publish <= now();