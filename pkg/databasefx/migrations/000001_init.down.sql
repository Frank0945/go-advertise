-- reverse: create "advertisement" table
DROP TABLE "public"."advertisement";

-- reverse: create function to check daily ad count
DROP FUNCTION check_daily_ad_count(value INTEGER);

-- reverse: create "daily_ad_count" table
DROP TABLE "public"."daily_ad_count";

-- reverse: create function to check active ads count
DROP FUNCTION check_active_ads_count();

-- reverse: create trigger to check active ads count
DROP TRIGGER check_active_ads_trigger;
