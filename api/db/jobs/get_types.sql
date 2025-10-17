SELECT jt.*
FROM job_types jt
WHERE jt.id IN (
    SELECT DISTINCT job_type_id
    FROM jobs
    WHERE 
        time_expire > NOW() AND
        time_publish < NOW() AND
        visible = TRUE
);