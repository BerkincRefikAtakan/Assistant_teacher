CREATE TABLE "teachers" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "surname" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "classes" (
  "id" bigserial PRIMARY KEY,
  "teacher_id" bigint NOT NULL,
  "class_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "students" (
  "id" bigserial PRIMARY KEY,
  "class_id" bigint NOT NULL,
  "name" varchar NOT NULL,
  "surname" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "paragraphs" (
  "id" bigserial PRIMARY KEY,
  "teacher_id" bigint NOT NULL,
  "header" varchar,
  "paragraph" varchar,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE INDEX ON "classes" ("teacher_id");

CREATE INDEX ON "students" ("class_id");

CREATE INDEX ON "paragraphs" ("header");

ALTER TABLE "classes" ADD FOREIGN KEY ("teacher_id") REFERENCES "teachers" ("id");

ALTER TABLE "students" ADD FOREIGN KEY ("class_id") REFERENCES "classes" ("id");

ALTER TABLE "paragraphs" ADD FOREIGN KEY ("teacher_id") REFERENCES "teachers" ("id");
