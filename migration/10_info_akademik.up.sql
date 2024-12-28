CREATE TABLE academic_information (
    id SERIAL PRIMARY KEY,
    student_id INTEGER REFERENCES student(id),
    details TEXT
);