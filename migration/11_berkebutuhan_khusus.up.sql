CREATE TABLE special_needs (
    id SERIAL PRIMARY KEY,
    student_id INTEGER REFERENCES student(id),
    detail INTEGER
);