SELECT DISTINCT ja.*
FROM jobs as j
LEFT JOIN job_types AS ja ON j.job_type = ja.id;