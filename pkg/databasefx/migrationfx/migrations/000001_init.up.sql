-- create "advertisements" table
CREATE TABLE "public"."advertisements" (
    "id" SERIAL NOT NULL,
    "title" VARCHAR(255),
    "start_at" TIMESTAMP,
    "end_at" TIMESTAMP,
    "age_end" INT NULL,
    "age_start" INT NULL,
    "country" VARCHAR(2) NULL,
    "gender" CHAR(1) NULL,
    "platform" VARCHAR(10) NULL,
    PRIMARY KEY ("id")
);