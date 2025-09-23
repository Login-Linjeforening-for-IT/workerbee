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
    "name_en" varchar NOT NULL,
    "description_no" varchar NOT NULL,
    "description_en" varchar NOT NULL,
    "informational_no" varchar NOT NULL,
    "informational_en" varchar NOT NULL,
    "time_type" time_type_enum NOT NULL DEFAULT 'default',
    "time_start" timestamp NOT NULL,
    "time_end" timestamp NOT NULL,
    "time_publish" timestamp,
    "time_signup_release" timestamp,
    "time_signup_deadline" timestamp,
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
    "color" varchar NOT NULL,
    "name_no" varchar NOT NULL,
    "name_en" varchar NOT NULL,
    "description_no" text NOT NULL,
    "description_en" text NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "audiences" (
    "id" SERIAL PRIMARY KEY,
    "name_no" varchar NOT NULL,
    "name_en" varchar NOT NULL,
    "description_no" varchar NOT NULL,
    "description_en" varchar NOT NULL,
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
    "shortname" varchar,
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
    "coordinate_long" float,
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
    "job_type" job_type NOT NULL DEFAULT 'full',
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
ALTER TABLE "events" ADD FOREIGN KEY ("parent_id") REFERENCES "events" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE "events" ADD FOREIGN KEY ("audience_id") REFERENCES "audiences" ("id");

ALTER TABLE "jobs" ADD FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id");
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


------------------
-- Dummy Data
------------------

INSERT INTO "skills" ("name") VALUES
('Programmering'),
('Datavisualisering'),
('Prosjektledelse'),
('Markedsføring'),
('Kundebehandling');

INSERT INTO "cities" ("name") VALUES
('Arendal'),
('Bodø'),
('Drammen'),
('Fredrikstad'),
('Gjøvik'),
('Halden'),
('Hamar'),
('Harstad'),
('Haugesund'),
('Kristiansand'),
('Kristiansund'),
('Larvik'),
('Lillehammer'),
('Mo i Rana'),
('Molde'),
('Moss'),
('Porsgrunn'),
('Sandefjord'),
('Sandnes'),
('Sarpsborg'),
('Skien'),
('Steinkjer'),
('Tønsberg'),
('Ålesund'),
('Alta'),
('Oslo'),
('Trondheim'),
('Bergen'),
('Stavanger'),
('Tromsø');


INSERT INTO "categories" (
  "color", 
  "name_no", 
  "name_en", 
  "description_no", 
  "description_en", 
  "updated_at", 
  "created_at"
)
VALUES
('#FF5733', 'Kultur', 'Culture', 'Kulturelle arrangementer og aktiviteter.', 'Cultural events and activities.', now(), now()),
('#33FF57', 'Teknologi', 'Technology', 'Arrangementer relatert til teknologi og innovasjon.', 'Events related to technology and innovation.', now(), now()),
('#3357FF', 'Sport', 'Sports', 'Sportslige aktiviteter og konkurranser.', 'Sports activities and competitions.', now(), now()),
('#FF33A8', 'Musikk', 'Music', 'Konserter og musikalske arrangementer.', 'Concerts and musical events.', now(), now()),
('#FFD700', 'Mat', 'Food', 'Arrangementer med fokus på mat og drikke.', 'Events focused on food and beverages.', now(), now());

INSERT INTO "locations" (
  "name_no", 
  "name_en", 
  "type", 
  "mazemap_campus_id", 
  "mazemap_poi_id", 
  "address_street", 
  "address_postcode", 
  "city_id", 
  "coordinate_lat", 
  "coordinate_long", 
  "url", 
  "updated_at", 
  "created_at"
)
VALUES
('Universitetet i Oslo', 'University of Oslo', 'address', NULL, NULL, 
  'Problemveien 7', 313, 26, 59.9390, 10.7205, NULL, now(), now()),
('Nidarosdomen', 'Nidaros Cathedral', 'coords', NULL, NULL, 
  NULL, NULL, 27, 63.4277, 10.3969, NULL, now(), now()),
('Bergenhus Festning', 'Bergenhus Fortress', 'address', NULL, NULL, 
  'Bergenhus', 5003, 28, 60.3993, 5.3221, NULL, now(), now()),
('Stavanger Forum', 'Stavanger Forum', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 1', 'Bestegata 1', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 2', 'Bestegata 2', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 3', 'Bestegata 3', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 4', 'Bestegata 4', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 5', 'Bestegata 5', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 6', 'Bestegata 6', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 7', 'Bestegata 7', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 8', 'Bestegata 8', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 9', 'Bestegata 9', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 10', 'Bestegata 10', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 11', 'Bestegata 11', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 12', 'Bestegata 12', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 13', 'Bestegata 13', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 14', 'Bestegata 14', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 15', 'Bestegata 15', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 16', 'Bestegata 16', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 17', 'Bestegata 17', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 18', 'Bestegata 18', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 19', 'Bestegata 19', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Bestegata 20', 'Bestegata 20', 'mazemap', 1, 12345, 
  NULL, 4021, 29, 58.9537, 5.6990, 'https://stavanger-forum.no', now(), now()),
('Tromsø Bibliotek', 'Tromsø Library', 'address', NULL, NULL, 
  'Grønnegata 94', 9008, 30, 69.6489, 18.9551, NULL, now(), now());
INSERT INTO "rules" (
  "name_no", "name_en", "description_no", "description_en", "updated_at", "created_at"
)
VALUES
('Ingen eksterne matvarer', 'No External Food Allowed', 
  'Det er ikke tillatt å ta med matvarer utenfra til arrangementet.', 
  'Bringing external food to the event is not allowed.', 
  now(), now()),
('Påmelding er obligatorisk', 'Registration Required', 
  'Alle deltakere må registrere seg på forhånd for å delta.', 
  'All participants must register in advance to attend.', 
  now(), now()),
('Stille mobiltelefoner', 'Silence Mobile Phones', 
  'Alle mobiltelefoner må settes på lydløs under arrangementet.', 
  'All mobile phones must be set to silent mode during the event.', 
  now(), now()),
('Ikke røyk', 'No Smoking', 
  'Røyking er ikke tillatt i eller rundt arrangementsområdet.', 
  'Smoking is not permitted in or around the event area.', 
  now(), now()),
('Gyldig billett kreves', 'Valid Ticket Required', 
  'Alle deltakere må ha en gyldig billett for å delta på arrangementet.', 
  'All attendees must present a valid ticket to participate in the event.', 
  now(), now());

INSERT INTO "audiences" (
  "name_no", "name_en", "description_no", "description_en", "updated_at", "created_at"
)
VALUES
('Studenter', 'Students', 
  'Arrangementet er rettet mot studenter fra alle studieretninger.', 
  'The event is targeted at students from all fields of study.', 
  now(), now()),
('Familier', 'Families', 
  'Arrangementet er familievennlig og åpent for alle aldersgrupper.', 
  'The event is family-friendly and open to all age groups.', 
  now(), now()),
('Profesjonelle', 'Professionals', 
  'Arrangementet er rettet mot fagpersoner innenfor relevante bransjer.', 
  'The event is aimed at professionals in relevant industries.', 
  now(), now());

INSERT INTO "organizations" (
  "shortname", 
  "name_no", 
  "name_en", 
  "description_no", 
  "description_en", 
  "type", 
  "link_homepage", 
  "link_linkedin", 
  "link_facebook", 
  "link_instagram", 
  "logo", 
  "updated_at", 
  "created_at"
)
VALUES
  ('UiO', 'Universitetet i Oslo', 'University of Oslo', 
   'Universitetet i Oslo er Norges største universitet, med et bredt fagtilbud.', 
   'The University of Oslo is Norways largest university, offering a wide range of programs.', 
   1, 'https://www.uio.no', 'https://www.linkedin.com/school/university-of-oslo', 
   'https://www.facebook.com/uni.oslo', 'https://www.instagram.com/uniofoslo', 
   'logo.png', now(), now()),
  ('NTNU', 'Norges teknisk-naturvitenskapelige universitet', 'Norwegian University of Science and Technology', 
   'NTNU er et teknisk universitet i Trondheim, kjent for sin forskning på teknologi og naturvitenskap.', 
   'NTNU is a technical university in Trondheim, known for its research in technology and natural sciences.', 
   2, 'https://www.ntnu.no', 'https://www.linkedin.com/school/ntnu', 
   'https://www.facebook.com/NTNU.no', 'https://www.instagram.com/ntnu_official', 
   'logo.png', now(), now()),
  ('DNB', 'DNB ASA', 'DNB ASA', 
   'DNB er Norges største finanskonsern med et bredt tilbud av finansielle tjenester.', 
   'DNB is Norways largest financial group, offering a wide range of financial services.', 
   3, 'https://www.dnb.no', 'https://www.linkedin.com/company/dnb', 
   'https://www.facebook.com/dnb.no', 'https://www.instagram.com/dnb.no', 
   'logo.png', now(), now()),
  ('Telenor', 'Telenor ASA', 'Telenor ASA', 
   'Telenor er et ledende teleselskap som tilbyr mobil- og bredbåndstjenester.', 
   'Telenor is a leading telecommunications company offering mobile and broadband services.', 
   3, 'https://www.telenor.no', 'https://www.linkedin.com/company/telenor', 
   'https://www.facebook.com/telenor', 'https://www.instagram.com/telenor', 
   'logo.png', now(), now()),
  ('SINTEF', 'SINTEF', 'SINTEF', 
   'SINTEF er en av Europas største uavhengige forskningsorganisasjoner, kjent for sitt arbeid innen teknologi og innovasjon.', 
   'SINTEF is one of Europes largest independent research organizations, known for its work in technology and innovation.', 
   4, 'https://www.sintef.no', 'https://www.linkedin.com/company/sintef', 
   'https://www.facebook.com/SINTEF', 'https://www.instagram.com/sintef', 
   'logo.png', now(), now());

INSERT INTO "jobs" (
  "visible", 
  "highlight", 
  "title_no", 
  "title_en", 
  "position_title_no", 
  "position_title_en", 
  "description_short_no", 
  "description_short_en", 
  "description_long_no", 
  "description_long_en", 
  "job_type", 
  "time_publish", 
  "time_expire", 
  "application_deadline", 
  "banner_image", 
  "organization_id", 
  "application_url", 
  "updated_at", 
  "created_at"
)
VALUES
(true, false, 'Softwareutvikler', 'Software Developer', 'Junior utvikler', 
  'Junior Developer', 'En spennende mulighet for nyutdannede til å utvikle programvare.', 
  'An exciting opportunity for recent graduates to develop software.', 
  'Som Junior Software Developer vil du være med på utvikling av applikasjoner og programvare.', 
  'As a Junior Software Developer, you will be involved in the development of applications and software.', 
  'full', now(), '2025-03-31', '2025-03-01', 'https://www.example.com/banner.jpg', 
  1, 'https://www.uio.no/job-apply', now(), now()),
(true, true, 'Markedsføringskoordinator', 'Marketing Coordinator', 'Markedsføringsspesialist', 
  'Marketing Specialist', 'Bli en del av vårt markedsføringsteam og jobb med spennende prosjekter.', 
  'Join our marketing team and work on exciting projects.', 
  'Som markedsføringskoordinator vil du ha ansvar for markedsføring og kommunikasjon på tvers av kanaler.', 
  'As a Marketing Coordinator, you will be responsible for marketing and communication across channels.', 
  'part', now(), '2025-05-31', '2025-04-15', 'https://www.example.com/banner2.jpg', 
  3, 'https://www.dnb.no/job-apply', now(), now()),
(false, false, 'Prosjektleder', 'Project Manager', 'Senior prosjektleder', 
  'Senior Project Manager', 'Vi søker en erfaren prosjektleder til å lede store prosjekter.', 
  'We are looking for an experienced project manager to lead large projects.', 
  'Som prosjektleder vil du ha ansvar for å lede prosjekter fra start til slutt, inkludert budsjett og tidsplanlegging.', 
  'As a Project Manager, you will be responsible for leading projects from start to finish, including budgeting and scheduling.', 
  'full', now(), '2027-06-30', '2027-05-01', 'https://www.example.com/banner3.jpg', 
  4, 'https://www.telenor.no/job-apply', now(), now()),
(true, false, 'Kundestøtteagent', 'Customer Support Agent', 'Kundestøtteansvarlig', 
  'Customer Support Manager', 'Bli en del av vårt kundeserviceteam og hjelp kunder med deres henvendelser.', 
  'Join our customer service team and assist customers with their inquiries.', 
  'Som kundestøtteansvarlig vil du hjelpe kunder via telefon, e-post og chat, samt sikre god kundetilfredshet.', 
  'As a Customer Support Manager, you will assist customers via phone, email, and chat, ensuring high customer satisfaction.', 
  'part', now(), '2027-04-15', '2027-03-15', 'https://www.example.com/banner4.jpg', 
  5, 'https://www.sintef.no/job-apply', now(), now()),
(true, true, 'Dataanalytiker', 'Data Analyst', 'Dataanalytiker', 
  'Data Analyst', 'Er du en dataanalytiker som elsker å finne innsikt fra store datamengder?', 
  'Are you a data analyst who loves to derive insights from large datasets?', 
  'Som dataanalytiker vil du analysere data for å identifisere trender og lage rapporter som støtter beslutningstaking.', 
  'As a Data Analyst, you will analyze data to identify trends and create reports that support decision-making.', 
  'full', now(), '2025-07-31', '2025-06-01', 'https://www.example.com/banner5.jpg', 
  2, 'https://www.ntnu.no/job-apply', now(), now());

INSERT INTO "ad_skill_relation" ("job_id", "skill_id")
VALUES
(1, 1), -- Software Developer (UiO) - Programmering
(1, 2), -- Software Developer (UiO) - Datavisualisering
(2, 3), -- Marketing Coordinator (DNB) - Prosjektledelse
(3, 4), -- Project Manager (Telenor) - Markedsføring
(3, 5); -- Project Manager (Telenor) - Kundebehandling

INSERT INTO "ad_city_relation" ("job_id", "city_id")
VALUES
(1, 26),  -- Software Developer (UiO) - Oslo
(2, 29),  -- Marketing Coordinator (DNB) - Stavanger
(3, 30),  -- Project Manager (Telenor) - Tromsø
(4, 27),  -- Customer Support Agent (SINTEF) - Trondheim
(5, 27);  -- Data Analyst (NTNU) - Trondheim

INSERT INTO "events" (
  "visible", "name_no", "name_en", "description_no", "description_en", 
  "informational_no", "informational_en", "time_type", "time_start", "time_end", 
  "time_publish", "time_signup_release", "time_signup_deadline", "canceled", 
  "digital", "highlight", "image_small", "image_banner", "link_facebook", 
  "link_discord", "link_signup", "link_stream", "capacity", "full", "category_id", 
  "organization_id", "location_id", "parent_id", "rule_id", "audience_id", "created_at", "updated_at"
)
VALUES
(true, 'Hackathon Oslo', 'Hackathon Oslo', 'Bli med på en spennende hackathon i Oslo!',
  'Join an exciting hackathon in Oslo!', 'Mer informasjon kommer snart.', 'More information coming soon.',
  'whole_day', '2025-02-01 09:00:00', '2025-02-01 18:00:00', '2025-02-01 09:00:00', 
  '2025-01-15 08:00:00', '2025-01-30 23:59:00', false, true, false, NULL, 'https://www.example.com/banner_hackathon.jpg', 
  NULL, NULL, NULL, NULL, NULL, false, 1, 1, 1, NULL, 1, 1, '2025-02-01 09:00:00', '2025-02-01 09:00:00'),
(true, 'Tech Conference Bergen', 'Tech Conference Bergen', 'Lær om de nyeste teknologiene på Tech Conference i Bergen.',
  'Learn about the latest technologies at Tech Conference in Bergen.', 'Påmelding nødvendig.', 'Registration required.',
  'whole_day', '2025-03-10 09:00:00', '2025-03-10 17:00:00', now(), 
  '2025-02-15 08:00:00', '2025-03-01 23:59:00', false, true, true, NULL, 'https://www.example.com/banner_tech_conference.jpg', 
  NULL, NULL, NULL, NULL, NULL, false, 2, 2, 2, NULL, 1, 2, now(), now()),
(true, 'AI Workshop Trondheim', 'AI Workshop Trondheim', 
  'Utforsk kunstig intelligens i Trondheim!', 
  'Explore artificial intelligence in Trondheim!', 
  'Gratis workshop for alle interesserte.', 'Free workshop for all interested.',
  'whole_day', NOW() + (INTERVAL '1 day' * trunc(random() * 30)), 
  NOW() + (INTERVAL '1 day' * trunc(random() * 30)) + (INTERVAL '1 hour' * trunc(random() * 10)), 
  NOW(), NOW() - INTERVAL '10 days', NOW() + INTERVAL '15 days', 
  false, true, false, NULL, 'https://www.example.com/banner_ai_workshop.jpg', 
  NULL, NULL, NULL, NULL, NULL, false, 3, 3, 3, NULL, 1, 1, NOW(), NOW()),
(true, 'Cybersecurity Summit Stavanger', 'Cybersecurity Summit Stavanger', 
  'Lær om cybersikkerhet i Stavanger!', 
  'Learn about cybersecurity in Stavanger!', 
  'Fokus på praktiske løsninger.', 'Focus on practical solutions.',
  'whole_day', NOW() + (INTERVAL '1 day' * trunc(random() * 30)), 
  NOW() + (INTERVAL '1 day' * trunc(random() * 30)) + (INTERVAL '1 hour' * trunc(random() * 10)), 
  NOW(), NOW() - INTERVAL '5 days', NOW() + INTERVAL '20 days', 
  true, true, true, NULL, 'https://www.example.com/banner_cybersecurity_summit.jpg', 
  NULL, NULL, NULL, NULL, NULL, false, 4, 3, 4, NULL, 3, 3, NOW(), NOW()),
(true, 'Cloud Computing Meetup Tromsø', 'Cloud Computing Meetup Tromsø', 
  'Møt eksperter innen skyteknologi i Tromsø.', 
  'Meet cloud technology experts in Tromsø.', 
  'Networking muligheter.', 'Networking opportunities.',
  'whole_day', NOW() - INTERVAL '5 days', 
  NOW() - INTERVAL '5 days' + (INTERVAL '1 hour' * trunc(random() * 10)), 
  NOW(), NOW() - INTERVAL '30 days', NOW() - INTERVAL '10 days', 
  false, true, false, NULL, 'https://www.example.com/banner_cloud_meetup.jpg', 
  NULL, NULL, NULL, NULL, NULL, false, 5, 1, 5, NULL, 2, 2, NOW(), NOW());

-- BeeFormed Dummy Data: users, forms, questions, options, submissions, answers, answer_options
-- 10 submissions per form (one per user), and 10 answers per question.
INSERT INTO users (full_name, email) VALUES
('User One','user1@example.com'),
('User Two','user2@example.com'),
('User Three','user3@example.com'),
('User Four','user4@example.com'),
('User Five','user5@example.com'),
('User Six','user6@example.com'),
('User Seven','user7@example.com'),
('User Eight','user8@example.com'),
('User Nine','user9@example.com'),
('User Ten','user10@example.com');

-- Two forms owned by user 1 and user 2
INSERT INTO forms (user_id, title, description, capacity, open_at, close_at)
VALUES
((SELECT id FROM users WHERE full_name='User One'), 'Form A', 'Test form A', 100, now(), '2026-12-31'),
((SELECT id FROM users WHERE full_name='User Two'), 'Form B', 'Test form B', 100, now(), '2026-12-31');

-- Questions for Form A (positions 1..5) using all types
INSERT INTO questions (form_id, question_title, question_description, question_type, required, position, max)
VALUES
((SELECT id FROM forms WHERE title='Form A'), 'A - Q1 Single choice', 'Single choice question', 'single_choice', true, 1, NULL),
((SELECT id FROM forms WHERE title='Form A'), 'A - Q2 Multiple choice', 'Multiple choice question', 'multiple_choice', true, 2, NULL),
((SELECT id FROM forms WHERE title='Form A'), 'A - Q3 Text', 'Open text question', 'text', false, 3, NULL),
((SELECT id FROM forms WHERE title='Form A'), 'A - Q4 Number', 'Numeric answer', 'number', false, 4, NULL),
((SELECT id FROM forms WHERE title='Form A'), 'A - Q5 Date', 'Date answer', 'date', false, 5, NULL);

-- Questions for Form B (positions 1..5) using all types
INSERT INTO questions (form_id, question_title, question_description, question_type, required, position, max)
VALUES
((SELECT id FROM forms WHERE title='Form B'), 'B - Q1 Single choice', 'Single choice question', 'single_choice', true, 1, NULL),
((SELECT id FROM forms WHERE title='Form B'), 'B - Q2 Multiple choice', 'Multiple choice question', 'multiple_choice', true, 2, NULL),
((SELECT id FROM forms WHERE title='Form B'), 'B - Q3 Text', 'Open text question', 'text', false, 3, NULL),
((SELECT id FROM forms WHERE title='Form B'), 'B - Q4 Number', 'Numeric answer', 'number', false, 4, NULL),
((SELECT id FROM forms WHERE title='Form B'), 'B - Q5 Date', 'Date answer', 'date', false, 5, NULL);

-- Options for single_choice and multiple_choice questions (4 options each)
-- Form A: single (pos=1) and multiple (pos=2)
INSERT INTO question_options (question_id, option_text, position)
SELECT q_single.id, opts.opt_text, opts.opt_pos
FROM (
  SELECT id FROM questions WHERE form_id=(SELECT id FROM forms WHERE title='Form A') AND position=1
) q_single,
(VALUES ('Option A1',1),('Option A2',2),('Option A3',3),('Option A4',4)) AS opts(opt_text,opt_pos)
ORDER BY q_single.id, opts.opt_pos;

INSERT INTO question_options (question_id, option_text, position)
SELECT q_multi.id, opts.opt_text, opts.opt_pos
FROM (
  SELECT id FROM questions WHERE form_id=(SELECT id FROM forms WHERE title='Form A') AND position=2
) q_multi,
(VALUES ('Multi A1',1),('Multi A2',2),('Multi A3',3),('Multi A4',4)) AS opts(opt_text,opt_pos)
ORDER BY q_multi.id, opts.opt_pos;

-- Form B: single (pos=1) and multiple (pos=2)
INSERT INTO question_options (question_id, option_text, position)
SELECT q_single.id, opts.opt_text, opts.opt_pos
FROM (
  SELECT id FROM questions WHERE form_id=(SELECT id FROM forms WHERE title='Form B') AND position=1
) q_single,
(VALUES ('Option B1',1),('Option B2',2),('Option B3',3),('Option B4',4)) AS opts(opt_text,opt_pos)
ORDER BY q_single.id, opts.opt_pos;

INSERT INTO question_options (question_id, option_text, position)
SELECT q_multi.id, opts.opt_text, opts.opt_pos
FROM (
  SELECT id FROM questions WHERE form_id=(SELECT id FROM forms WHERE title='Form B') AND position=2
) q_multi,
(VALUES ('Multi B1',1),('Multi B2',2),('Multi B3',3),('Multi B4',4)) AS opts(opt_text,opt_pos)
ORDER BY q_multi.id, opts.opt_pos;

-- Create 10 submissions per form (one per user)
INSERT INTO submissions (form_id, user_id, submitted_at)
SELECT f.id, u.id, now()
FROM users u CROSS JOIN forms f
WHERE f.title='Form A';

INSERT INTO submissions (form_id, user_id, submitted_at)
SELECT f.id, u.id, now()
FROM users u CROSS JOIN forms f
WHERE f.title='Form B';

-- Answers:
-- 1) single_choice: one answer per submission (10 answers per single_choice question)
INSERT INTO answers (submission_id, question_id, option_id, created_at, updated_at)
SELECT s.id, q.id,
  (SELECT id FROM question_options qo WHERE qo.question_id = q.id AND qo.position = ((s.user_id-1)%4)+1 LIMIT 1),
  now(), now()
FROM submissions s
JOIN questions q ON q.form_id = s.form_id
WHERE q.question_type = 'single_choice';

-- 2) multiple_choice: create one answers row per submission (so 10 answers per multiple_choice question),
--    and store chosen options in answer_options (two options per submission)
INSERT INTO answers (submission_id, question_id, created_at, updated_at)
SELECT s.id, q.id, now(), now()
FROM submissions s
JOIN questions q ON q.form_id = s.form_id
WHERE q.question_type = 'multiple_choice';

-- populate answer_options: pick two option positions per submission (pos and pos+1 wrap)
INSERT INTO answer_options (answer_id, option_id)
SELECT a.id,
       qo.id
FROM answers a
JOIN submissions s ON s.id = a.submission_id
JOIN questions q ON q.id = a.question_id
JOIN question_options qo ON qo.question_id = q.id
WHERE q.question_type = 'multiple_choice'
  AND qo.position IN ( ((s.user_id-1)%4)+1, ((s.user_id)%4)+1 );

-- 3) text answers (one per submission -> 10 answers per text question)
INSERT INTO answers (submission_id, question_id, answer_text, created_at, updated_at)
SELECT s.id, q.id, 'Text answer from user ' || s.user_id, now(), now()
FROM submissions s
JOIN questions q ON q.form_id = s.form_id
WHERE q.question_type = 'text';

-- 4) number answers (store as text in answer_text)
INSERT INTO answers (submission_id, question_id, answer_text, created_at, updated_at)
SELECT s.id, q.id, (s.user_id * 10)::text, now(), now()
FROM submissions s
JOIN questions q ON q.form_id = s.form_id
WHERE q.question_type = 'number';

-- 5) date answers (store as ISO date string in answer_text)
INSERT INTO answers (submission_id, question_id, answer_text, created_at, updated_at)
SELECT s.id, q.id, to_char((CURRENT_DATE + (s.user_id || ' days')::interval)::date, 'YYYY-MM-DD'), now(), now()
FROM submissions s
JOIN questions q ON q.form_id = s.form_id
WHERE q.question_type = 'date';