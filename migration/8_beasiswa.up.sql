CREATE TABLE scholarship (
    id SERIAL PRIMARY KEY,
    student_id INTEGER REFERENCES student(id),
    details TEXT
);