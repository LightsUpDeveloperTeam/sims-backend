CREATE TYPE document_type AS ENUM ('commission', 'internship_report');

CREATE TABLE internship_document (
    id SERIAL PRIMARY KEY,
    student_id INT NOT NULL REFERENCES student(id),
    document document_type DEFAULT 'internship_report' NOT NULL,
    document_file VARCHAR NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
