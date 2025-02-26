CREATE TABLE internship_contract (
    id SERIAL PRIMARY KEY,
    student_id INT NOT NULL REFERENCES student(id),
    company_name VARCHAR UNIQUE NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    contract_description TEXT NOT NULL,
    deleted_by BIGINT,
    deleted_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
