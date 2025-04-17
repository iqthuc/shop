CREATE TABLE "categories" ("id" serial PRIMARY KEY, "name" varchar NOT NULL);

CREATE TABLE "brands" (
    "id" serial PRIMARY KEY,
    "name" varchar UNIQUE NOT NULL
);

CREATE TABLE "products" (
    "id" serial PRIMARY KEY,
    "name" varchar NOT NULL,
    "slug" varchar UNIQUE NOT NULL,
    "desciprtion" text,
    "category_id" int,
    "brand_id" int,
    "main_image_url" varchar,
    "base_price" decimal(10, 2) DEFAULT 0
);

CREATE TABLE "attributes" (
    "id" serial PRIMARY KEY,
    "name" varchar UNIQUE NOT NULL
);

CREATE TABLE "attribute_values" (
    "id" serial PRIMARY KEY,
    "attribute_id" int,
    "value" varchar
);

CREATE TABLE "product_variants" (
    "id" serial PRIMARY KEY,
    "product_id" int,
    "sku" varchar UNIQUE NOT NULL,
    "price" decimal(15, 2) NOT NULL,
    "stock_quantity" int NOT NULL DEFAULT 0,
    "sold" int NOT NULL DEFAULT 0,
    "image_url" varchar,
    "is_default" bool NOT NULL DEFAULT false
);

CREATE TABLE "variant_attribute_values" ("variant_id" int, "value_id" int);

CREATE INDEX ON "products" ("name");

CREATE INDEX ON "products" ("category_id");

CREATE INDEX ON "products" ("brand_id");

CREATE UNIQUE INDEX ON "attribute_values" ("attribute_id", "value");

CREATE INDEX ON "product_variants" ("product_id");

CREATE INDEX ON "product_variants" ("sku");

CREATE INDEX ON "product_variants" ("price");

CREATE INDEX ON "product_variants" ("stock_quantity");

CREATE UNIQUE INDEX ON "variant_attribute_values" ("variant_id", "value_id");

COMMENT ON COLUMN "attributes"."name" IS 'eg: color, size';

COMMENT ON COLUMN "attribute_values"."value" IS 'eg: red, blue, S, M';

ALTER TABLE "products" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("brand_id") REFERENCES "brands" ("id");

ALTER TABLE "attribute_values" ADD FOREIGN KEY ("attribute_id") REFERENCES "attributes" ("id");

ALTER TABLE "product_variants" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "variant_attribute_values" ADD FOREIGN KEY ("variant_id") REFERENCES "product_variants" ("id");

ALTER TABLE "variant_attribute_values" ADD FOREIGN KEY ("value_id") REFERENCES "attribute_values" ("id");
