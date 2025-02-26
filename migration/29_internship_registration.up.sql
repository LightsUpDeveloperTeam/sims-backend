CREATE TABLE internship_registration (
    id SERIAL PRIMARY KEY,
    student_id INT NOT NULL REFERENCES student(id),
    internship_vacancy_id INT NOT NULL REFERENCES internship_vacancy(id),
    company_name VARCHAR NOT NULL UNIQUE,
    position VARCHAR NOT NULL,
    registration_date DATE NOT NULL,
    deleted_by BIGINT,
    deleted_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);