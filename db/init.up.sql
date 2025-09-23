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
    "highlight" bool NOT NULL DEFAULT false,
    "name_no" text NOT NULL,
    "name_en" text NOT NULL,
    "informational_no" text NOT NULL,
    "informational_en" text NOT NULL,
    "description_no" text NOT NULL,
    "description_en" text NOT NULL,
    "time_type" time_type_enum NOT NULL DEFAULT 'default',
    "time_start" timestamp NOT NULL,
    "time_end" timestamp NOT NULL,
    "time_publish" timestamp,
    "time_signup_release" timestamp,
    "time_signup_deadline" timestamp,
    "link_signup" text,
    "capacity" int,
    "full" bool NOT NULL DEFAULT false,
    "canceled" bool NOT NULL DEFAULT false,
    "digital" bool NOT NULL DEFAULT false,
    "image_small" text,
    "image_banner" text NOT NULL,
    "link_facebook" text,
    "link_discord" text,
    "link_stream" text,
    "category_id" int NOT NULL,
    "organization_id" int,
    "location_id" int,
    "parent_id" int,
    "rule_id" int,
    "audience_id" int,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE "categories" (
    "id" SERIAL PRIMARY KEY,
    "name_no" text NOT NULL,
    "name_en" text NOT NULL,
    "description_no" text NOT NULL,
    "description_en" text NOT NULL,
    "color" text NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "audiences" (
    "id" SERIAL PRIMARY KEY,
    "name_no" text NOT NULL,
    "name_en" text NOT NULL,
    "description_no" text NOT NULL,
    "description_en" text NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "jobs" (
    "id" SERIAL PRIMARY KEY,
    "visible" bool NOT NULL DEFAULT false,
    "highlight" bool NOT NULL DEFAULT false,
    "title_no" text NOT NULL,
    "title_en" text NOT NULL,
    "position_title_no" text NOT NULL,
    "position_title_en" text NOT NULL,
    "description_short_no" text NOT NULL,
    "description_short_en" text NOT NULL,
    "description_long_no" text NOT NULL,
    "description_long_en" text NOT NULL,
    "job_type" job_type NOT NULL DEFAULT 'full',
    "time_publish" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "time_expire" timestamp NOT NULL,
    "application_url" text,
    "banner_image" text,
    "organization_id" int NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "ad_city_relation" (
    "job_id" int NOT NULL,
    "city_id" int NOT NULL,
    PRIMARY KEY ("job_id", "city_id")
);

CREATE TABLE "cities" (
    "id" SERIAL PRIMARY KEY,
    "name" text UNIQUE
);

CREATE TABLE "ad_skill_relation" (
    "job_id" int NOT NULL,
    "skill_id" int NOT NULL,
    PRIMARY KEY ("job_id", "skill_id")
);

CREATE TABLE "skills" (
    "id" SERIAL PRIMARY KEY,
    "name" text NOT NULL
);

CREATE TABLE "organizations" (
    "id" SERIAL PRIMARY KEY,
    "name_no" text NOT NULL,
    "name_en" text NOT NULL,
    "description_no" text NOT NULL,
    "description_en" text NOT NULL,
    "type" int NOT NULL DEFAULT 1,
    "link_homepage" text,
    "link_linkedin" text,
    "link_facebook" text,
    "link_instagram" text,
    "logo" text,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "locations" (
    "id" SERIAL PRIMARY KEY,
    "name_no" text NOT NULL,
    "name_en" text NOT NULL,
    "type" location_type NOT NULL DEFAULT 'digital',
    "mazemap_campus_id" int,
    "mazemap_poi_id" int,
    "address_street" text,
    "address_postcode" int,
    "city_id" int,
    "coordinate_lat" float,
    "coordinate_long" float,
    "url" text,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "rules" (
    "id" SERIAL PRIMARY KEY,
    "name_no" text NOT NULL,
    "name_en" text NOT NULL,
    "description_no" text NOT NULL,
    "description_en" text NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX ON "events" ("visible");
CREATE INDEX ON "events" ("highlight");
CREATE INDEX ON "events" ("category_id");
CREATE INDEX ON "events" ("time_start");
CREATE INDEX ON "events" ("time_end");
CREATE INDEX ON "events" ("updated_at");
CREATE INDEX ON "events" ("created_at");

CREATE INDEX ON "categories" ("updated_at");
CREATE INDEX ON "categories" ("created_at");

CREATE INDEX ON "audiences" ("updated_at");
CREATE INDEX ON "audiences" ("created_at");

CREATE INDEX ON "rules" ("updated_at");
CREATE INDEX ON "rules" ("created_at");

CREATE INDEX ON "organizations" ("type");
CREATE INDEX ON "organizations" ("updated_at");
CREATE INDEX ON "organizations" ("created_at");

CREATE INDEX ON "locations" ("updated_at");
CREATE INDEX ON "locations" ("created_at");

CREATE INDEX ON "jobs" ("updated_at");
CREATE INDEX ON "jobs" ("created_at");
CREATE INDEX ON "ad_city_relation" ("job_id");
CREATE INDEX ON "ad_city_relation" ("city_id");
CREATE INDEX ON "ad_skill_relation" ("job_id");
CREATE INDEX ON "ad_skill_relation" ("skill_id");

ALTER TABLE "events" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");
ALTER TABLE "events" ADD FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id");
ALTER TABLE "events" ADD FOREIGN KEY ("location_id") REFERENCES "locations" ("id");
ALTER TABLE "events" ADD FOREIGN KEY ("rule_id") REFERENCES "rules" ("id");
ALTER TABLE "events" ADD FOREIGN KEY ("parent_id") REFERENCES "events" ("id");
ALTER TABLE "events" ADD FOREIGN KEY ("audience_id") REFERENCES "audiences" ("id");

ALTER TABLE "jobs" ADD FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id");
ALTER TABLE "ad_city_relation" ADD FOREIGN KEY ("job_id") REFERENCES "jobs" ("id");
ALTER TABLE "ad_city_relation" ADD FOREIGN KEY ("city_id") REFERENCES "cities" ("id");
ALTER TABLE "ad_skill_relation" ADD FOREIGN KEY ("job_id") REFERENCES "jobs" ("id");
ALTER TABLE "ad_skill_relation" ADD FOREIGN KEY ("skill_id") REFERENCES "skills" ("id");

ALTER TABLE "locations" ADD FOREIGN KEY ("city_id") REFERENCES "cities" ("id");

-- BeeFormed
CREATE TYPE question_type_enum AS ENUM (
    'single_choice',
    'multiple_choice',
    'text',
    'number',
    'date'
);

CREATE TABLE users (
    "id" SERIAL PRIMARY KEY,
    "full_name" text UNIQUE NOT NULL,
    "email" text UNIQUE NOT NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE forms (
    "id" SERIAL PRIMARY KEY,
    "user_id" int NOT NULL REFERENCES "users"("id") ON DELETE CASCADE,
    "title" text NOT NULL,
    "description" text NOT NULL,
    "capacity" int,
    "open_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "close_at" timestamp NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE questions (
  "id" SERIAL PRIMARY KEY,
  "form_id" int NOT NULL REFERENCES "forms"("id") ON DELETE CASCADE,
    "question_title" text NOT NULL,
    "question_description" text NOT NULL,
    "question_type" question_type_enum NOT NULL,
    "required" boolean DEFAULT false,
    "position" int NOT NULL,
    "max" int,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE question_options (
    "id" SERIAL PRIMARY KEY,
    "question_id" int NOT NULL REFERENCES "questions"("id") ON DELETE CASCADE,
    "option_text" text NOT NULL,
    "position" int NOT NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE submissions (
    "id" SERIAL PRIMARY KEY,
    "form_id" int NOT NULL REFERENCES "forms"("id") ON DELETE CASCADE,
    "user_id" int NOT NULL REFERENCES "users"("id") ON DELETE CASCADE,
    "submitted_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE answers (
    "id" SERIAL PRIMARY KEY,
    "submission_id" int NOT NULL REFERENCES "submissions"("id") ON DELETE CASCADE,
    "question_id" int NOT NULL REFERENCES "questions"("id") ON DELETE CASCADE,
    "option_id" int REFERENCES "question_options"("id"),
    "answer_text" text,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE answer_options (
    "answer_id" int NOT NULL REFERENCES "answers"("id") ON DELETE CASCADE,
    "option_id" int NOT NULL REFERENCES "question_options"("id") ON DELETE CASCADE,
    PRIMARY KEY ("answer_id", "option_id")
);

CREATE INDEX ON "forms"("user_id");

CREATE INDEX ON "questions"("form_id");
CREATE INDEX ON "question_options"("question_id");

CREATE INDEX ON "submissions"("form_id");
CREATE INDEX ON "submissions"("user_id");

CREATE INDEX ON "answers"("submission_id");
CREATE INDEX ON "answers"("question_id");
CREATE INDEX ON "answers"("option_id");
