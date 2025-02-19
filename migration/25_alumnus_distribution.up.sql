CREATE TABLE alumnus_distribution (
    id SERIAL PRIMARY KEY,
    student_id INT NOT NULL REFERENCES student(id),
    company_name VARCHAR NOT NULL UNIQUE,
    position VARCHAR NOT NULL,
    start_date DATE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);