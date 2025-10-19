SELECT s.* 
FROM skills s
WHERE 
    s.id IN (SELECT skill_id FROM ad_skill_relation WHERE job_id IN 
        (SELECT id FROM jobs WHERE visible = true AND time_expire > NOW() AND time_publish < NOW())
    );