SELECT s.* 
FROM skills s
WHERE 
    s.id IN (SELECT skill_id FROM ad_skill_relation);