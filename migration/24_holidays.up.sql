CREATE TABLE holidays (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    type VARCHAR NOT NULL, -- ('religious', 'national', 'international') NOT NULL,
    date DATE NOT NULL,
    is_holiday BOOLEAN NOT NULL,
    is_celebrated_at_school BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);