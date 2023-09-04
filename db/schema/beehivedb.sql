-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2023-05-04T12:50:58.754Z

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

CREATE TABLE "Event" (
  "id" SERIAL PRIMARY KEY,
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
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Category" (
  "id" SERIAL PRIMARY KEY,
  "color" char(6) NOT NULL,
  "name_no" varchar NOT NULL,
  "name_en" varchar,
  "description_no" text NOT NULL,
  "description_en" text,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Audience" (
  "id" SERIAL PRIMARY KEY,
  "name_no" varchar NOT NULL,
  "name_en" varchar,
  "description_no" varchar NOT NULL,
  "description_en" varchar,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "EventAudienceRelation" (
  "event" int NOT NULL,
  "audience" int NOT NULL,
  PRIMARY KEY ("event", "audience")
);

CREATE TABLE "Rule" (
  "id" SERIAL PRIMARY KEY,
  "name_no" varchar NOT NULL,
  "name_en" varchar,
  "description_no" varchar NOT NULL,
  "description_en" varchar,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Organization" (
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
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "EventOrganizationRelation" (
  "event" int NOT NULL,
  "organization" varchar NOT NULL,
  PRIMARY KEY ("event", "organization")
);

CREATE TABLE "Location" (
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
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "JobAdvertisement" (
  "id" SERIAL PRIMARY KEY,
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
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "AdCityRelation" (
  "ad" int NOT NULL,
  "city" varchar NOT NULL
);

CREATE TABLE "City" (
  "name" varchar PRIMARY KEY
);

CREATE TABLE "Skill" (
  "ad" int NOT NULL,
  "skill" varchar NOT NULL,
  PRIMARY KEY ("ad", "skill")
);

CREATE INDEX ON "Event" ("highlight");

CREATE INDEX ON "Event" ("category");

CREATE INDEX ON "Event" ("updated_at");

CREATE INDEX ON "Event" ("created_at");

CREATE INDEX ON "Category" ("updated_at");

CREATE INDEX ON "Category" ("created_at");

CREATE INDEX ON "Audience" ("updated_at");

CREATE INDEX ON "Audience" ("created_at");

CREATE INDEX ON "EventAudienceRelation" ("event");

CREATE INDEX ON "EventAudienceRelation" ("audience");

CREATE INDEX ON "Rule" ("updated_at");

CREATE INDEX ON "Rule" ("created_at");

CREATE INDEX ON "Organization" ("updated_at");

CREATE INDEX ON "Organization" ("created_at");

CREATE INDEX ON "EventOrganizationRelation" ("event");

CREATE INDEX ON "EventOrganizationRelation" ("organization");

CREATE INDEX ON "Location" ("updated_at");

CREATE INDEX ON "Location" ("created_at");

CREATE INDEX ON "JobAdvertisement" ("updated_at");

CREATE INDEX ON "JobAdvertisement" ("created_at");

CREATE INDEX ON "AdCityRelation" ("ad");

CREATE INDEX ON "AdCityRelation" ("city");

CREATE INDEX ON "AdCityRelation" ("ad", "city");

CREATE INDEX ON "Skill" ("ad");

CREATE INDEX ON "Skill" ("skill");

COMMENT ON COLUMN "Category"."color" IS 'hex color';

ALTER TABLE "Event" ADD FOREIGN KEY ("category") REFERENCES "Category" ("id");

ALTER TABLE "Event" ADD FOREIGN KEY ("location") REFERENCES "Location" ("id");

ALTER TABLE "Event" ADD FOREIGN KEY ("parent") REFERENCES "Event" ("id");

ALTER TABLE "Event" ADD FOREIGN KEY ("rule") REFERENCES "Rule" ("id");

ALTER TABLE "EventAudienceRelation" ADD FOREIGN KEY ("event") REFERENCES "Event" ("id");

ALTER TABLE "EventAudienceRelation" ADD FOREIGN KEY ("audience") REFERENCES "Audience" ("id");

ALTER TABLE "EventOrganizationRelation" ADD FOREIGN KEY ("event") REFERENCES "Event" ("id");

ALTER TABLE "EventOrganizationRelation" ADD FOREIGN KEY ("organization") REFERENCES "Organization" ("shortname");

ALTER TABLE "JobAdvertisement" ADD FOREIGN KEY ("organization") REFERENCES "Organization" ("shortname");

ALTER TABLE "AdCityRelation" ADD FOREIGN KEY ("ad") REFERENCES "JobAdvertisement" ("id");

ALTER TABLE "AdCityRelation" ADD FOREIGN KEY ("city") REFERENCES "City" ("name");

ALTER TABLE "Skill" ADD FOREIGN KEY ("ad") REFERENCES "JobAdvertisement" ("id");
