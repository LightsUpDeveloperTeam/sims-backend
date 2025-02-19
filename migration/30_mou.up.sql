CREATE TABLE mou (
    id SERIAL PRIMARY KEY,
    industry_partner_id INT NOT NULL REFERENCES industry_partner(id),
    mou_number VARCHAR NOT NULL UNIQUE,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    description TEXT NOT NULL,
    mou_file VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
