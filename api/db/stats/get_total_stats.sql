-- name: GetTotalStats :one
SELECT
  (SELECT COUNT(*) FROM event) AS total_events,
  (SELECT COUNT(*) FROM job) AS total_jobs,
  (SELECT COUNT(*) FROM organization) AS total_organizations,
  (SELECT COUNT(*) FROM location) AS total_locations,
  (SELECT COUNT(*) FROM rule) AS total_rules;