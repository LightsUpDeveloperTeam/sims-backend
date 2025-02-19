CREATE TABLE industry_partner (
    id SERIAL PRIMARY KEY,
    partner_name VARCHAR NOT NULL UNIQUE,
    work_sector VARCHAR NOT NULL UNIQUE,
    address VARCHAR NOT NULL UNIQUE,
    email VARCHAR NOT NULL UNIQUE,
    telephone_number VARCHAR NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
