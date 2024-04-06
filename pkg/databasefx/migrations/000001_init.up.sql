-- create "advertisement" table
CREATE TABLE "public"."advertisement" (
    "id" SERIAL NOT NULL,
    "title" VARCHAR(255),
    "start_at" TIMESTAMP,
    "end_at" TIMESTAMP,
    "age_end" INT NULL,
    "age_start" INT NULL,
    "country" VARCHAR(2) [] NULL,
    "gender" CHAR(1) [] NULL,
    "platform" VARCHAR(10) [] NULL,
    PRIMARY KEY ("id")
);

-- create function to check daily ad count
CREATE OR REPLACE FUNCTION check_daily_ad_count(value INTEGER)
RETURNS BOOLEAN AS $$
BEGIN
    IF value > 3000 THEN
        RAISE EXCEPTION 'Daily ads count exceeded';
    END IF;
    RETURN TRUE;
END;
$$ LANGUAGE plpgsql;

-- create "daily_ad_count" table
CREATE TABLE "public"."daily_ad_count" (
    "date" DATE NOT NULL,
    "count" INT NOT NULL CHECK (check_daily_ad_count("count")),
    PRIMARY KEY ("date")
);

-- create function to check active ads count
CREATE OR REPLACE FUNCTION check_active_ads_count()
RETURNS TRIGGER AS $$
DECLARE
    active_ads_count INTEGER;
BEGIN
    SELECT COUNT(*) INTO active_ads_count 
    FROM "advertisement"
    WHERE "start_at" < NOW() AND "end_at" > NOW();

    IF active_ads_count >= 1000 THEN
        RAISE EXCEPTION 'Active ads count exceeded';
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- create trigger to check active ads count
CREATE TRIGGER check_active_ads_trigger
BEFORE INSERT ON "advertisement"
FOR EACH ROW
EXECUTE FUNCTION check_active_ads_count();
