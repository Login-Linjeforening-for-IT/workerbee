CREATE TYPE "time_type_enum" AS ENUM (
  'default',
  'no_end',
  'whole_day',
  'tbd'
);

CREATE TYPE "location_type" AS ENUM (
  'mazemap',
  'coords',
  'address',
  'digital'
);

CREATE TYPE "job_type" AS ENUM (
  'full',
  'part',
  'summer',
  'verv'
);

CREATE TABLE "events" (
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
  "time_end" timestamptz NOT NULL,
  "time_publish" timestamptz,
  "time_signup_release" timestamptz,
  "time_signup_deadline" timestamptz,
  "canceled" bool NOT NULL DEFAULT false,
  "digital" bool NOT NULL DEFAULT false,
  "highlight" bool NOT NULL DEFAULT false,
  "image_small" varchar,
  "image_banner" varchar NOT NULL,
  "link_facebook" varchar,
  "link_discord" varchar,
  "link_signup" varchar,
  "link_stream" varchar,
  "capacity" int,
  "full" bool NOT NULL DEFAULT false,
  "category_id" int NOT NULL,
  "location_id" int,
  "parent_id" int,
  "rule_id" int,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "categories" (
  "id" SERIAL PRIMARY KEY,
  "color" varchar NOT NULL,
  "name_no" varchar NOT NULL,
  "name_en" varchar,
  "description_no" text NOT NULL,
  "description_en" text,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "audiences" (
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
  "event_id" int NOT NULL,
  "audience_id" int NOT NULL,
  PRIMARY KEY ("event_id", "audience_id")
);

CREATE TABLE "rules" (
  "id" SERIAL PRIMARY KEY,
  "name_no" varchar NOT NULL,
  "name_en" varchar,
  "description_no" varchar NOT NULL,
  "description_en" varchar,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "organizations" (
  "id" SERIAL PRIMARY KEY,
  "shortname" varchar,
  "name_no" varchar NOT NULL,
  "name_en" varchar,
  "description_no" varchar NOT NULL,
  "description_en" varchar,
  "type" int NOT NULL DEFAULT 1,
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
  "event_id" int NOT NULL,
  "organization_id" int NOT NULL,
  PRIMARY KEY ("event_id", "organization_id")
);

CREATE TABLE "locations" (
  "id" SERIAL PRIMARY KEY,
  "name_no" varchar NOT NULL,
  "name_en" varchar,
  "type" location_type NOT NULL DEFAULT 'digital',
  "mazemap_campus_id" int,
  "mazemap_poi_id" int,
  "address_street" varchar,
  "address_postcode" int,
  "city_id" int,
  "coordinate_lat" float,
  "coordinate_long" float,
  "url" varchar,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "job_advertisements" (
  "id" SERIAL PRIMARY KEY,
  "visible" bool NOT NULL DEFAULT false,
  "highlight" bool NOT NULL DEFAULT false,
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
  "time_expire" timestamptz NOT NULL,
  "application_deadline" timestamptz NOT NULL,
  "banner_image" varchar,
  "organization_id" int NOT NULL,
  "application_url" varchar,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "ad_city_relation" (
  "job_advertisement_id" int NOT NULL,
  "city_id" int NOT NULL,
  PRIMARY KEY ("job_advertisement_id", "city_id")
);

CREATE TABLE "cities" (
  "id" SERIAL,
  "name" varchar,
  PRIMARY KEY ("id", "name")
);

CREATE TABLE "ad_skill_relation" (
  "job_advertisement_id" int NOT NULL,
  "skill_id" int NOT NULL,
  PRIMARY KEY ("job_advertisement_id", "skill_id")
);

CREATE TABLE "skills" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE INDEX ON "events" ("visible");

CREATE INDEX ON "events" ("highlight");

CREATE INDEX ON "events" ("category_id");

CREATE INDEX ON "events" ("time_start");

CREATE INDEX ON "events" ("time_end");

CREATE INDEX ON "events" ("updated_at");

CREATE INDEX ON "events" ("created_at");

CREATE INDEX ON "events" ("deleted_at");

CREATE INDEX ON "categories" ("updated_at");

CREATE INDEX ON "categories" ("created_at");

CREATE INDEX ON "audiences" ("updated_at");

CREATE INDEX ON "audiences" ("created_at");

CREATE INDEX ON "audiences" ("deleted_at");

CREATE INDEX ON "event_audience_relation" ("event_id");

CREATE INDEX ON "event_audience_relation" ("audience_id");

CREATE INDEX ON "rules" ("updated_at");

CREATE INDEX ON "rules" ("created_at");

CREATE INDEX ON "rules" ("deleted_at");

CREATE INDEX ON "organizations" ("type");

CREATE INDEX ON "organizations" ("updated_at");

CREATE INDEX ON "organizations" ("created_at");

CREATE INDEX ON "organizations" ("deleted_at");

CREATE INDEX ON "event_organization_relation" ("event_id");

CREATE INDEX ON "event_organization_relation" ("organization_id");

CREATE INDEX ON "locations" ("updated_at");

CREATE INDEX ON "locations" ("created_at");

CREATE INDEX ON "locations" ("deleted_at");

CREATE INDEX ON "job_advertisements" ("updated_at");

CREATE INDEX ON "job_advertisements" ("created_at");

CREATE INDEX ON "job_advertisements" ("deleted_at");

CREATE INDEX ON "ad_city_relation" ("job_advertisement_id");

CREATE INDEX ON "ad_city_relation" ("city_id");

CREATE INDEX ON "ad_skill_relation" ("job_advertisement_id");

CREATE INDEX ON "ad_skill_relation" ("skill_id");

ALTER TABLE "events" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "events" ADD FOREIGN KEY ("location_id") REFERENCES "locations" ("id");

ALTER TABLE "locations" ADD FOREIGN KEY ("city_id") REFERENCES "cities" ("id");

ALTER TABLE "events" ADD FOREIGN KEY ("parent_id") REFERENCES "events" ("id");

ALTER TABLE "events" ADD FOREIGN KEY ("rule_id") REFERENCES "rules" ("id");

ALTER TABLE "event_audience_relation" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

ALTER TABLE "event_audience_relation" ADD FOREIGN KEY ("audience_id") REFERENCES "audiences" ("id");

ALTER TABLE "event_organization_relation" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

ALTER TABLE "event_organization_relation" ADD FOREIGN KEY ("organization_id") REFERENCES "organizations" ("shortname");

ALTER TABLE "job_advertisements" ADD FOREIGN KEY ("organization_id") REFERENCES "organizations" ("shortname");

ALTER TABLE "ad_city_relation" ADD FOREIGN KEY ("job_advertisement_id") REFERENCES "job_advertisements" ("id");

ALTER TABLE "ad_city_relation" ADD FOREIGN KEY ("city_id") REFERENCES "cities" ("name");

ALTER TABLE "ad_skill_relation" ADD FOREIGN KEY ("job_advertisement_id") REFERENCES "job_advertisements" ("id");

ALTER TABLE "ad_skill_relation" ADD FOREIGN KEY ("skill_id") REFERENCES "skills" ("id");
