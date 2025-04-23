CREATE TABLE "carts" (
    "id" serial PRIMARY KEY,
    "user_id" uuid NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz
);

CREATE TABLE "cart_items" (
    "id" serial PRIMARY KEY,
    "cart_id" int NOT NULL,
    "product_variant_id" int NOT NULL,
    "quantity" int NOT NULL DEFAULT 1,
    "created_at" timestamptz NOT NULL DEFAULT (now ()),
    "updated_at" timestamptz NOT NULL DEFAULT (now ())
);

CREATE INDEX ON "carts" ("user_id");

ALTER TABLE "carts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE UNIQUE INDEX ON "cart_items" ("cart_id", "product_variant_id");

ALTER TABLE "cart_items" ADD FOREIGN KEY ("cart_id") REFERENCES "carts" ("id");

ALTER TABLE "cart_items" ADD FOREIGN KEY ("product_variant_id") REFERENCES "product_variants" ("id");
