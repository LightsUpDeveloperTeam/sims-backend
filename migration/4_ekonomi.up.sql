CREATE TABLE economy (
    id SERIAL PRIMARY KEY,
    family_id INTEGER REFERENCES family(id),
    housing_access INTEGER,
    kitchen_fuel INTEGER,
    electricity_capacity INTEGER,
    monthly_consumption_expenditure INTEGER,
    water_source INTEGER
);