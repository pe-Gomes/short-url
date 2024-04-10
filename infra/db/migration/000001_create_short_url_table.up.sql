CREATE TABLE IF NOT EXISTS "users" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "short_links" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigserial,
  "url" text NOT NULL,
  "slug" text NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE INDEX ON "users" ("id");

CREATE INDEX ON "short_links" ("id");

CREATE INDEX ON "short_links" ("url");

ALTER TABLE "short_links" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");