CREATE TABLE "users" (
    "id" uuid PRIMARY KEY,
    "username" varchar(32) UNIQUE,
    "email" varchar(128) UNIQUE NOT NULL,
    "hashed_password" (64) varchar,
    "role" varchar(16) NOT NULL DEFAULT 'user',
    "created_at" timestamptz NOT NULL DEFAULT (now ()),
    "updated_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE INDEX ON "users" ("email");
