CREATE TABLE education (
    id SERIAL PRIMARY KEY,
    student_id INTEGER REFERENCES student(id),
    class VARCHAR,
    nis VARCHAR,
    nisn INTEGER,
    level VARCHAR,
    aspiration INTEGER,
    hobby INTEGER
);