SELECT array_agg(e.enumlabel ORDER BY e.enumsortorder) AS job_type_no_values
FROM pg_type t
JOIN pg_enum e ON t.oid = e.enumtypid
WHERE t.typname = 'job_type_no';