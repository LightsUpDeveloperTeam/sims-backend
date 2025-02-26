CREATE TABLE internship_vacancy (
    id SERIAL PRIMARY KEY,
    industry_partner_id INT NOT NULL REFERENCES industry_partner(id),
    position_name VARCHAR NOT NULL,
    description TEXT NOT NULL,
    open_date DATE NOT NULL,
    close_date DATE NOT NULL,
    deleted_by BIGINT,
    deleted_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);