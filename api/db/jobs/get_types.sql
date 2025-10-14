SELECT DISTINCT job_type FROM jobs
WHERE 
    time_expire > NOW() AND
    time_publish < NOW() AND
    canceled = FALSE AND
    visible = TRUE;