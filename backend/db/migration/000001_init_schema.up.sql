CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "fistname" varchar NOT NULL,
  "lastname" bigint NOT NULL,
  "email" varchar NOT NULL,
  "phno" varchar NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now()
);
