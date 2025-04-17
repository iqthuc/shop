CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    username VARCHAR(32) UNIQUE,
    email VARCHAR(128) UNIQUE NOT NULL,
    password_hash VARCHAR(64),
    role VARCHAR(16) DEFAULT 'user',
    created_at TIMESTAMP DEFAULT now (),
    updated_at TIMESTAMP
);
