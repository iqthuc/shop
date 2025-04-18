DROP TABLE variant_attribute_values;

DROP TABLE product_variants;

DROP TABLE attribute_values;

DROP TABLE products;

DROP TABLE attributes;

DROP TABLE brands;

DROP TABLE categories;

DROP INDEX IF EXISTS idx_product_name_trgm;

DROP EXTENSION IF EXISTS pg_trgm;
