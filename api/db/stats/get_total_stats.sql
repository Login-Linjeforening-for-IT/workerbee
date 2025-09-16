-- name: GetTotalStats :one
SELECT
  (SELECT COUNT(*) FROM events) AS total_events,
  (SELECT COUNT(*) FROM job_advertisements) AS total_jobs,
  (SELECT COUNT(*) FROM organizations) AS total_organizations,
  (SELECT COUNT(*) FROM locations) AS total_locations,
  (SELECT COUNT(*) FROM rules) AS total_rules;