CREATE TABLE address (
    id SERIAL PRIMARY KEY,
    family_id INTEGER REFERENCES family(id),
    address VARCHAR,
    hamlet VARCHAR,
    regency INTEGER,
    district INTEGER,
    village INTEGER,
    postal_code INTEGER,
    latitude FLOAT,
    longitude FLOAT,
    primary_address BOOLEAN,
    province INTEGER,
    rt INTEGER,
    rw INTEGER
);