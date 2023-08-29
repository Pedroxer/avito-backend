CREATE TABLE "segments" (
  "id" bigserial PRIMARY KEY,
  "name" text NOT NULL
);

CREATE TABLE "users" (
  "id" integer PRIMARY KEY
  
);

CREATE TABLE "seg_to_user" (
  "user_id" integer NOT NULL,
  "seg_id" integer NOT NULL 
);

ALTER TABLE "seg_to_user" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "seg_to_user" ADD FOREIGN KEY ("seg_id") REFERENCES "segments" ("id");
