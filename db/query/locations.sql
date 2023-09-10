-- name: GetLocation :one
SELECT * FROM "location" WHERE "id" = sqlc.arg('id')::int LIMIT 1;

-- name: GetMazemapLocations :many
SELECT "id", "name_no", "name_en", 
        "mazemap_campus_id", "mazemap_poi_id",
        "updated_at", "url", "deleted_at" IS NOT NULL AS "is_deleted"
    FROM "location"
    WHERE "type" = 'mazemap'
    LIMIT sqlc.arg('limit')::int
    OFFSET sqlc.arg('offset')::int;

-- name: GetAddressLocations :many
SELECT "id", "name_no", "name_en", 
        "address_street", "address_postcode", "city_name",
        "updated_at", "url", "deleted_at" IS NOT NULL AS "is_deleted"
    FROM "location"
    WHERE "type" = 'address'
    LIMIT sqlc.arg('limit')::int
    OFFSET sqlc.arg('offset')::int;

-- name: GetCoordsLocations :many
SELECT "id", "name_no", "name_en", 
        "coordinate_lat", "coordinate_long",
        "updated_at", "url", "deleted_at" IS NOT NULL AS "is_deleted"
    FROM "location"
    WHERE "type" = 'coords'
    LIMIT sqlc.arg('limit')::int
    OFFSET sqlc.arg('offset')::int;

-- name: GetLocations :many
SELECT *, "deleted_at" IS NOT NULL AS "is_deleted" FROM "location"
    LIMIT sqlc.arg('limit')::int
    OFFSET sqlc.arg('offset')::int;

-- name: CreateLocation :one
INSERT INTO "location" (
    "name_no", "name_en",
    "address_street", "address_postcode", "city_name",
    "coordinate_lat", "coordinate_long",
    "mazemap_campus_id", "mazemap_poi_id",
    "type", "url"
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING *;

-- name: UpdateLocation :one
UPDATE "location"
SET
    "name_no" = COALESCE(sqlc.narg(name_no), name_no),
    "name_en" = COALESCE(sqlc.narg(name_en), name_en),
    "address_street" = COALESCE(sqlc.narg(address_street), address_street),
    "address_postcode" = COALESCE(sqlc.narg(address_postcode), address_postcode),
    "city_name" = COALESCE(sqlc.narg(city_name), city_name),
    "coordinate_lat" = COALESCE(sqlc.narg(coordinate_lat), coordinate_lat),
    "coordinate_long" = COALESCE(sqlc.narg(coordinate_long), coordinate_long),
    "mazemap_campus_id" = COALESCE(sqlc.narg(mazemap_campus_id), mazemap_campus_id),
    "mazemap_poi_id" = COALESCE(sqlc.narg(mazemap_poi_id), mazemap_poi_id),
    "type" = COALESCE(sqlc.narg(type), type),
    "url" = COALESCE(sqlc.narg(url), url),
    "updated_at" = now()
WHERE "id" = sqlc.arg(id)::int
RETURNING *;

-- name: SoftDeleteLocation :one
UPDATE "location"
SET
    "deleted_at" = now(),
    "updated_at" = now()
WHERE "id" = sqlc.arg('id')::int
RETURNING *;
