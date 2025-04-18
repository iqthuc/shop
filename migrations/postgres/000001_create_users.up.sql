CREATE TABLE "users" (
    "id" uuid PRIMARY KEY,
    "username" varchar(32) UNIQUE,
    "email" varchar(128) UNIQUE NOT NULL,
    "password_hash" varchar(64),
    "role" varchar(16) NOT NULL DEFAULT 'user',
    "created_at" timestamptz NOT NULL DEFAULT (now ()),
    "updated_at" timestamptz
);

CREATE INDEX ON "users" ("email");
