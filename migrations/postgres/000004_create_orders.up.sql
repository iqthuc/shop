CREATE TYPE "OrderStatus" AS ENUM (
    'Pending',
    'Processing',
    'Shipped',
    'Delivered',
    'Cancelled'
);

CREATE TYPE "PaymenStatus" AS ENUM ('Pending', 'Paid', 'Failed');

CREATE TABLE "orders" (
    "id" serial PRIMARY KEY,
    "user_id" uuid NOT NULL,
    "product_variant_id" int NOT NULL,
    "status" OrderStatus NOT NULL DEFAULT 'Pending',
    "total_amount" decimal(15, 2) NOT NULL,
    "payment_status" PaymenStatus NOT NULL DEFAULT 'Pending',
    "created_at" timestamptz NOT NULL DEFAULT (now ()),
    "updated_at" timestamptz NOT NULL DEFAULT (now ())
);

CREATE INDEX ON "orders" ("user_id");

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("product_variant_id") REFERENCES "product_variants" ("id");

CREATE TABLE "order_items" (
    "id" serial PRIMARY KEY,
    "order_id" int NOT NULL,
    "product_variant_id" int NOT NULL,
    "quantity" int NOT NULL,
    "price" demimal (15, 2) NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now ()),
    "updated_at" timestamptz NOT NULL DEFAULT (now ())
);

CREATE INDEX ON "order_items" ("order_id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("product_variant_id") REFERENCES "product_variants" ("id");
