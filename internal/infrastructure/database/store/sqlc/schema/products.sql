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
