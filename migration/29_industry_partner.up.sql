CREATE TABLE industry_partner (
    id SERIAL PRIMARY KEY,
    partner_name VARCHAR NOT NULL UNIQUE,
    field_of_work jsonb NOT NULL UNIQUE,
    address VARCHAR NOT NULL UNIQUE,
    email VARCHAR NOT NULL UNIQUE,
    telephone_number VARCHAR NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
