CREATE TYPE "time_type_enum" AS ENUM (
    'default',
    'no_end',
    'whole_day',
    'to_be_determined'
);

CREATE TYPE "location_type" AS ENUM (
    'mazemap',
    'coords',
    'address',
    'digital'
);

CREATE TYPE "job_type" AS ENUM (
    'full_time',
    'part_time',
    'summer',
    'verv'
);

CREATE TYPE "job_type_no" AS ENUM (
    'full_tid',
    'del_tid',
    'sommer',
    'verv'
);

CREATE TYPE categories AS ENUM (
  'TekKom',
  'CTFKom',
  'EvntKom',
  'PR',
  'Social',
  'Login',
  'Buddyweek',
  'BedKom',
  'Careerdays',
  'Cyberdays',
  'Other'
);

CREATE TYPE categories_no AS ENUM (
  'TekKom',
  'CTFKom',
  'Evntkom',
  'PR',
  'Sosialt',
  'Login',
  'Fadderuka',
  'BedKom',
  'Karrieredagene',
  'Cyberdagene',
  'Other'
);

CREATE TYPE audience AS ENUM (
  'students',
  'first_semester',
  'second_semester',
  'third_semester',
  'fourth_semester',
  'fifth_semester',
  'sixth_semester',
  'seventh_semester',
  'Login',
  'open',
  'bachelor',
  'master',
  'phd'
);

CREATE TYPE audience_no AS ENUM (
  'students',
  'første_semester',
  'andre_semester',
  'tredje_semester',
  'fjerde_semester',
  'femte_semester',
  'sjette_semester',
  'sjuende_semester',
  'Login',
  'åpen',
  'bachelor',
  'master',
  'phd'
);

CREATE TABLE "events" (
    "id" SERIAL PRIMARY KEY,
    "visible" bool NOT NULL DEFAULT false,
    "name_no" varchar NOT NULL,
    "name_en" varchar NOT NULL,
    "description_no" varchar NOT NULL,
    "description_en" varchar NOT NULL,
    "informational_no" varchar NOT NULL,
    "informational_en" varchar NOT NULL,
    "time_type" time_type_enum NOT NULL DEFAULT 'default',
    "time_start" timestamp NOT NULL,
    "time_end" timestamp NOT NULL,
    "time_publish" timestamp NOT NULL,
    "time_signup_release" timestamp,
    "time_signup_deadline" timestamp,
    "canceled" bool NOT NULL DEFAULT false,
    "category" categories NOT NULL,
    "digital" bool NOT NULL DEFAULT false,
    "highlight" bool NOT NULL DEFAULT false,
    "image_small" varchar,
    "image_banner" varchar,
    "link_facebook" varchar,
    "link_discord" varchar,
    "link_signup" varchar,
    "link_stream" varchar,
    "capacity" int,
    "full_time" bool NOT NULL DEFAULT false,
    "organization_id" int,
    "location_id" int,
    "parent_id" int,
    "rule_id" int,
    "audience" audience,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "rules" (
    "id" SERIAL PRIMARY KEY,
    "name_no" varchar NOT NULL,
    "name_en" varchar NOT NULL,
    "description_no" varchar NOT NULL,
    "description_en" varchar NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "organizations" (
    "id" SERIAL PRIMARY KEY,
    "name_no" varchar NOT NULL,
    "name_en" varchar NOT NULL,
    "description_no" varchar NOT NULL,
    "description_en" varchar NOT NULL,
    "link_homepage" varchar NOT NULL,
    "link_linkedin" varchar,
    "link_facebook" varchar,
    "link_instagram" varchar,
    "logo" varchar,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "locations" (
    "id" SERIAL PRIMARY KEY,
    "name_no" varchar NOT NULL,
    "name_en" varchar NOT NULL,
    "type" location_type NOT NULL DEFAULT 'digital',
    "mazemap_campus_id" int,
    "mazemap_poi_id" int,
    "address_street" varchar,
    "address_postcode" int,
    "city_id" int,
    "coordinate_lat" float,
    "coordinate_lon" float,
    "url" varchar,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "jobs" (
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
    "job_type" job_type NOT NULL DEFAULT 'full_time',
    "time_publish" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "time_expire" timestamp NOT NULL,
    "application_deadline" timestamp NOT NULL,
    "banner_image" varchar,
    "organization_id" int NOT NULL,
    "application_url" varchar,
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
    "name" varchar UNIQUE
);

CREATE TABLE "ad_skill_relation" (
    "job_id" int NOT NULL,
    "skill_id" int NOT NULL,
    PRIMARY KEY ("job_id", "skill_id")
);

CREATE TABLE "skills" (
    "id" SERIAL PRIMARY KEY,
    "name" varchar NOT NULL
);

CREATE INDEX ON "events" ("visible");
CREATE INDEX ON "events" ("highlight");
CREATE INDEX ON "events" ("time_start");
CREATE INDEX ON "events" ("time_end");
CREATE INDEX ON "events" ("updated_at");
CREATE INDEX ON "events" ("created_at");


CREATE INDEX ON "rules" ("updated_at");
CREATE INDEX ON "rules" ("created_at");

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

ALTER TABLE "events" ADD FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE "events" ADD FOREIGN KEY ("location_id") REFERENCES "locations" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE "events" ADD FOREIGN KEY ("rule_id") REFERENCES "rules" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE "events" ADD FOREIGN KEY ("parent_id") REFERENCES "events" ("id") ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE "jobs" ADD FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE "ad_city_relation" ADD FOREIGN KEY ("job_id") REFERENCES "jobs" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE "ad_city_relation" ADD FOREIGN KEY ("city_id") REFERENCES "cities" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE "ad_skill_relation" ADD FOREIGN KEY ("job_id") REFERENCES "jobs" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE "ad_skill_relation" ADD FOREIGN KEY ("skill_id") REFERENCES "skills" ("id") ON UPDATE CASCADE ON DELETE CASCADE;

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
    "full_name" varchar UNIQUE NOT NULL,
    "email" varchar UNIQUE NOT NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE forms (
    "id" SERIAL PRIMARY KEY,
    "user_id" int NOT NULL REFERENCES "users"("id") ON DELETE CASCADE,
    "title" varchar NOT NULL,
    "description" varchar NOT NULL,
    "capacity" int,
    "open_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "close_at" timestamp NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE questions (
  "id" SERIAL PRIMARY KEY,
  "form_id" int NOT NULL REFERENCES "forms"("id") ON DELETE CASCADE,
    "question_title" varchar NOT NULL,
    "question_description" varchar NOT NULL,
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
    "option_text" varchar NOT NULL,
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