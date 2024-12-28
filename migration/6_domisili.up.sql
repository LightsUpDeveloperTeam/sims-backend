CREATE TABLE residency (
    id SERIAL PRIMARY KEY,
    student_id INTEGER REFERENCES student(id),
    residence_distance VARCHAR,
    residence_status INTEGER,
    transportation INTEGER,
    travel_time VARCHAR
);