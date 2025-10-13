SELECT
    ja.*,
    cities,
    skills,
    COUNT(*) OVER() AS total_count
FROM jobs ja
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
WHERE (
    $1 = '' OR to_json(ja)::text ILIKE '%' || $1 || '%'
    )
    AND (
        cardinality($2::text[]) = 0
        OR ja.job_type = ANY($2::job_type[])
    )
    AND (
        cardinality($3::text[]) = 0
        OR skill_agg.skills::text[] && $3::text[]
    )
    AND (
        cardinality($4::text[]) = 0
        OR city_agg.cities::text[] && $4::text[]
    )
