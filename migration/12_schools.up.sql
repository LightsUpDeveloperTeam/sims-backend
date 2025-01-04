DROP TYPE IF EXISTS subscription_status_type CASCADE;

CREATE TYPE subscription_status_type AS ENUM ('active', 'inactive', 'trial', 'expired');

CREATE TABLE IF NOT EXISTS schools (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    address TEXT,
    contact_email VARCHAR(255) NOT NULL UNIQUE,
    contact_phone VARCHAR(50),
    logo_url VARCHAR(255),
    subscription_status subscription_status_type DEFAULT 'trial',
    subscription_plan VARCHAR(255),
    subscription_expiry_date DATE,
    deleted_by BIGINT,
    deleted_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
