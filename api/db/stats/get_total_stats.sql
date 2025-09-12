-- name: GetTotalStats :one
SELECT
  (SELECT COUNT(*) FROM event WHERE deleted_at IS NULL) AS total_events,
  (SELECT COUNT(*) FROM job WHERE deleted_at IS NULL) AS total_jobs,
  (SELECT COUNT(*) FROM organization WHERE deleted_at IS NULL) AS total_organizations,
  (SELECT COUNT(*) FROM location WHERE deleted_at IS NULL) AS total_locations,
  (SELECT COUNT(*) FROM rule WHERE deleted_at IS NULL) AS total_rules;