DROP TYPE IF EXISTS status_collaboration CASCADE;

CREATE TYPE status_collaboration AS ENUM ('Active', 'Inactive');

CREATE TABLE collaboration_history (
    id SERIAL PRIMARY KEY,
    industry_partner_id INT NOT NULL REFERENCES industry_partner(id),
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    description TEXT NOT NULL,
    status status_collaboration DEFAULT 'Inactive' NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
