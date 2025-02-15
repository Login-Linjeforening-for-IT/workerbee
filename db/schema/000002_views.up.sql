------------------
-- Events
------------------

-- Events that are visible to the public
--   Not deleted and is visible
CREATE OR REPLACE VIEW "public_event" AS
SELECT * FROM "event"
WHERE "deleted_at" IS NULL
    AND "visible" IS TRUE
    AND ("time_publish" IS NULL OR "time_publish" < now());

-- Public events that are upcoming
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
