CREATE TABLE "segments" (
    "id" bigserial PRIMARY KEY,
    "name" text
);

CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "name" text UNIQUE NOT NULL 
);

CREATE TABLE "seg_to_user" (
    "user_id" integer,
    "seg_id" integer
);

ALTER TABLE "seg_to_user" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "seg_to_user" ADD FOREIGN KEY ("seg_id") REFERENCES "segments" ("id");