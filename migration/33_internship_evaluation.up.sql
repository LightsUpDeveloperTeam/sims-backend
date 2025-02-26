DROP TYPE IF EXISTS evaluation_rating CASCADE;

CREATE TYPE evaluation_rating AS ENUM ('Good', 'Bad');

CREATE TABLE internship_evaluation (
    id SERIAL PRIMARY KEY,
    student_id INT NOT NULL REFERENCES student(id),
    date DATE NOT NULL,
    rating evaluation_rating DEFAULT 'Bad' NOT NULL,
    deleted_by BIGINT,
    deleted_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);