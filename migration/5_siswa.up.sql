CREATE TABLE student (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    birth_place VARCHAR,
    birth_date TIMESTAMP,
    gender BOOLEAN,
    religion INTEGER,
    child_order INTEGER,
    number_of_siblings INTEGER,
    citizenship VARCHAR,
    email VARCHAR,
    family_id INTEGER REFERENCES family(id)
);