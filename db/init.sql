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
  'city',
  'digital'
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
  "name_en" varchar NOT NULL,
  "description_no" varchar NOT NULL,
  "description_en" varchar NOT NULL,
  "informational_no" varchar NOT NULL,
  "informational_en" varchar NOT NULL,
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
  "category" int NOT NULL,
  "location" int,
  "parent" int,
  "rule" int,
  "audience" int,
  "organization" int,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "category" (
  "id" SERIAL PRIMARY KEY,
  "color" varchar NOT NULL,
  "name_no" varchar NOT NULL,
  "name_en" varchar NOT NULL,
  "description_no" text NOT NULL,
  "description_en" text NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "audience" (
  "id" SERIAL PRIMARY KEY,
  "name_no" varchar NOT NULL,
  "name_en" varchar NOT NULL,
  "description_no" varchar NOT NULL,
  "description_en" varchar NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "rule" (
  "id" SERIAL PRIMARY KEY,
  "name_no" varchar NOT NULL,
  "name_en" varchar NOT NULL,
  "description_no" varchar NOT NULL,
  "description_en" varchar NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "organization" (
  "id" SERIAL PRIMARY KEY,
  "shortname" varchar NOT NULL,
  "name_no" varchar NOT NULL,
  "name_en" varchar NOT NULL,
  "description_no" varchar NOT NULL,
  "description_en" varchar NOT NULL,
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

CREATE TABLE "location" (
  "id" SERIAL PRIMARY KEY,
  "name_no" varchar NOT NULL,
  "name_en" varchar NOT NULL,
  "type" location_type NOT NULL DEFAULT 'digital',
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
  "highlight" bool NOT NULL DEFAULT false,
  "title_no" varchar NOT NULL,
  "title_en" varchar NOT NULL,
  "position_title_no" varchar NOT NULL,
  "position_title_en" varchar NOT NULL,
  "description_short_no" varchar NOT NULL,
  "description_short_en" varchar NOT NULL,
  "description_long_no" varchar NOT NULL,
  "description_long_en" varchar NOT NULL,
  "job_type" job_type NOT NULL DEFAULT 'full',
  "time_publish" timestamptz NOT NULL DEFAULT (now()),
  "time_expire" timestamptz NOT NULL,
  "application_deadline" timestamptz NOT NULL,
  "banner_image" varchar,
  "organization" int NOT NULL,
  "application_url" varchar,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "job_skill_relation" (
  "job_id" int NOT NULL,
  "skill" varchar NOT NULL,
  PRIMARY KEY ("job_id", "skill")
);

CREATE TABLE "job_location_relation" (
  "job_id" int NOT NULL,
  "location_id" int NOT NULL,
  PRIMARY KEY ("job_id", "location_id")
);

CREATE INDEX ON "event" ("visible");

CREATE INDEX ON "event" ("category");

CREATE INDEX ON "event" ("time_start");

CREATE INDEX ON "event" ("time_end");

CREATE INDEX ON "event" ("deleted_at");

CREATE INDEX ON "audience" ("deleted_at");

CREATE INDEX ON "rule" ("deleted_at");

CREATE INDEX ON "organization" ("type");

CREATE INDEX ON "organization" ("deleted_at");

CREATE INDEX ON "location" ("deleted_at");

CREATE INDEX ON "job_advertisement" ("deleted_at");

CREATE INDEX ON "job_location_relation" ("job_id");

CREATE INDEX ON "job_location_relation" ("location_id");

CREATE INDEX ON "job_skill_relation" ("job_id");

CREATE INDEX ON "job_skill_relation" ("skill");

ALTER TABLE "event" ADD FOREIGN KEY ("category") REFERENCES "category" ("id");

ALTER TABLE "event" ADD FOREIGN KEY ("location") REFERENCES "location" ("id");

ALTER TABLE "event" ADD FOREIGN KEY ("parent") REFERENCES "event" ("id");

ALTER TABLE "event" ADD FOREIGN KEY ("rule") REFERENCES "rule" ("id");

ALTER TABLE "event" ADD FOREIGN KEY ("audience") REFERENCES "audience" ("id");

ALTER TABLE "event" ADD FOREIGN KEY ("organization") REFERENCES "organization" ("id");

ALTER TABLE "job_advertisement" ADD FOREIGN KEY ("organization") REFERENCES "organization" ("id");

ALTER TABLE "job_location_relation" ADD FOREIGN KEY ("job_id") REFERENCES "job_advertisement" ("id");

ALTER TABLE "job_location_relation" ADD FOREIGN KEY ("location_id") REFERENCES "location" ("id");

ALTER TABLE "job_skill_relation" ADD FOREIGN KEY ("job_id") REFERENCES "job_advertisement" ("id");


------------------
-- Events
------------------

CREATE OR REPLACE VIEW "public_event" AS
SELECT * FROM "event"
WHERE "deleted_at" IS NULL
    AND "visible" IS TRUE
    AND ("time_publish" IS NULL OR "time_publish" < now());

CREATE OR REPLACE VIEW "visible_event" AS
SELECT * FROM "public_event"
WHERE "time_end" > now();


------------------
-- Jobs
------------------

CREATE OR REPLACE VIEW "public_job" AS
SELECT * FROM "job_advertisement"
WHERE "deleted_at" IS NULL
    AND "visible" IS TRUE
    AND "time_publish" < now();

CREATE OR REPLACE VIEW "visible_job" AS
SELECT * FROM "public_job"
WHERE "time_expire" > now();
