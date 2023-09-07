-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2023-09-06T23:02:35.301Z

CREATE TYPE "time_type_enum" AS ENUM (
  'default',
  'no_end',
  'whole_day'
);

CREATE TYPE "location_type" AS ENUM (
  'mazemap',
  'coords',
  'address',
  'none'
);

CREATE TYPE "job_type" AS ENUM (
  'full',
  'part',
  'summer',
  'verv'
);

CREATE TABLE "event" (
  "id" SERIAL PRIMARY KEY,
  "visible" bool NOT NULL DEFAULT false,
  "name_no" varchar NOT NULL,
  "name_en" varchar,
  "description_no" varchar NOT NULL,
  "description_en" varchar,
  "informational_no" varchar,
  "informational_en" varchar,
  "time_type" time_type_enum NOT NULL DEFAULT 'default',
  "time_start" timestamptz NOT NULL,
  "time_end" timestamptz,
  "time_publish" timestamptz,
  "time_signup_release" timestamptz,
  "time_signup_deadline" timestamptz,
  "canceled" bool NOT NULL DEFAULT false,
  "digital" bool NOT NULL DEFAULT false,
  "highlight" bool NOT NULL DEFAULT false,
  "image_small" varchar NOT NULL,
  "image_banner" varchar NOT NULL,
  "link_facebook" varchar NOT NULL,
  "link_discord" varchar NOT NULL,
  "link_signup" varchar NOT NULL,
  "link_stream" varchar,
  "capacity" int,
  "full" bool NOT NULL DEFAULT false,
  "category" int NOT NULL,
  "location" int,
  "parent" int,
  "rule" int,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "category" (
  "id" SERIAL PRIMARY KEY,
  "color" char(6) NOT NULL,
  "name_no" varchar NOT NULL,
  "name_en" varchar,
  "description_no" text NOT NULL,
  "description_en" text,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "audience" (
  "id" SERIAL PRIMARY KEY,
  "name_no" varchar NOT NULL,
  "name_en" varchar,
  "description_no" varchar NOT NULL,
  "description_en" varchar,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "event_audience_relation" (
  "event" int NOT NULL,
  "audience" int NOT NULL,
  PRIMARY KEY ("event", "audience")
);

CREATE TABLE "rule" (
  "id" SERIAL PRIMARY KEY,
  "name_no" varchar NOT NULL,
  "name_en" varchar,
  "description_no" varchar NOT NULL,
  "description_en" varchar,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "organization" (
  "shortname" varchar PRIMARY KEY,
  "name_no" varchar NOT NULL,
  "name_en" varchar,
  "description_no" varchar NOT NULL,
  "description_en" varchar,
  "link_homepage" varchar,
  "link_linkedin" varchar,
  "link_facebook" varchar,
  "link_instagram" varchar,
  "logo" varchar,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "event_organization_relation" (
  "event" int NOT NULL,
  "organization" varchar NOT NULL,
  PRIMARY KEY ("event", "organization")
);

CREATE TABLE "location" (
  "id" SERIAL PRIMARY KEY,
  "name_no" varchar NOT NULL,
  "name_en" varchar,
  "type" location_type NOT NULL DEFAULT 'none',
  "mazemap_campus_id" int,
  "mazemap_poi_id" int,
  "address_street" varchar,
  "address_postcode" int,
  "city_name" varchar,
  "coordinate_lat" float,
  "coordinate_long" float,
  "url" varchar,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "job_advertisement" (
  "id" SERIAL PRIMARY KEY,
  "visible" bool NOT NULL DEFAULT false,
  "title_no" varchar NOT NULL,
  "title_en" varchar,
  "position_title_no" varchar NOT NULL,
  "position_title_en" varchar,
  "description_short_no" varchar NOT NULL,
  "description_short_en" varchar,
  "description_long_no" varchar NOT NULL,
  "description_long_en" varchar,
  "job_type" job_type NOT NULL DEFAULT 'full',
  "time_publish" timestamptz NOT NULL DEFAULT (now()),
  "application_deadline" timestamptz NOT NULL,
  "banner_image" varchar,
  "organization" varchar NOT NULL,
  "application_url" varchar NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "ad_city_relation" (
  "ad" int NOT NULL,
  "city" varchar NOT NULL,
  PRIMARY KEY ("ad", "city")
);

CREATE TABLE "city" (
  "name" varchar PRIMARY KEY
);

CREATE TABLE "skill" (
  "ad" int NOT NULL,
  "skill" varchar NOT NULL,
  PRIMARY KEY ("ad", "skill")
);

CREATE INDEX ON "event" ("visible");

CREATE INDEX ON "event" ("highlight");

CREATE INDEX ON "event" ("category");

CREATE INDEX ON "event" ("time_start");

CREATE INDEX ON "event" ("time_end");

CREATE INDEX ON "event" ("updated_at");

CREATE INDEX ON "event" ("created_at");

CREATE INDEX ON "event" ("deleted_at");

CREATE INDEX ON "category" ("updated_at");

CREATE INDEX ON "category" ("created_at");

CREATE INDEX ON "audience" ("updated_at");

CREATE INDEX ON "audience" ("created_at");

CREATE INDEX ON "audience" ("deleted_at");

CREATE INDEX ON "event_audience_relation" ("event");

CREATE INDEX ON "event_audience_relation" ("audience");

CREATE INDEX ON "rule" ("updated_at");

CREATE INDEX ON "rule" ("created_at");

CREATE INDEX ON "rule" ("deleted_at");

CREATE INDEX ON "organization" ("updated_at");

CREATE INDEX ON "organization" ("created_at");

CREATE INDEX ON "organization" ("deleted_at");

CREATE INDEX ON "event_organization_relation" ("event");

CREATE INDEX ON "event_organization_relation" ("organization");

CREATE INDEX ON "location" ("updated_at");

CREATE INDEX ON "location" ("created_at");

CREATE INDEX ON "location" ("deleted_at");

CREATE INDEX ON "job_advertisement" ("updated_at");

CREATE INDEX ON "job_advertisement" ("created_at");

CREATE INDEX ON "job_advertisement" ("deleted_at");

CREATE INDEX ON "ad_city_relation" ("ad");

CREATE INDEX ON "ad_city_relation" ("city");

CREATE INDEX ON "skill" ("ad");

CREATE INDEX ON "skill" ("skill");

COMMENT ON COLUMN "category"."color" IS 'hex color';

ALTER TABLE "event" ADD FOREIGN KEY ("category") REFERENCES "category" ("id");

ALTER TABLE "event" ADD FOREIGN KEY ("location") REFERENCES "location" ("id");

ALTER TABLE "event" ADD FOREIGN KEY ("parent") REFERENCES "event" ("id");

ALTER TABLE "event" ADD FOREIGN KEY ("rule") REFERENCES "rule" ("id");

ALTER TABLE "event_audience_relation" ADD FOREIGN KEY ("event") REFERENCES "event" ("id");

ALTER TABLE "event_audience_relation" ADD FOREIGN KEY ("audience") REFERENCES "audience" ("id");

ALTER TABLE "event_organization_relation" ADD FOREIGN KEY ("event") REFERENCES "event" ("id");

ALTER TABLE "event_organization_relation" ADD FOREIGN KEY ("organization") REFERENCES "organization" ("shortname");

ALTER TABLE "job_advertisement" ADD FOREIGN KEY ("organization") REFERENCES "organization" ("shortname");

ALTER TABLE "ad_city_relation" ADD FOREIGN KEY ("ad") REFERENCES "job_advertisement" ("id");

ALTER TABLE "ad_city_relation" ADD FOREIGN KEY ("city") REFERENCES "city" ("name");

ALTER TABLE "skill" ADD FOREIGN KEY ("ad") REFERENCES "job_advertisement" ("id");
