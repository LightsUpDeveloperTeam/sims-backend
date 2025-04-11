CREATE TYPE progress_status AS ENUM ('Open', 'OnGoing', 'Finish');

CREATE TABLE internship_progress (
    id SERIAL PRIMARY KEY,
    student_id INT NOT NULL REFERENCES student(id),
    date DATE NOT NULL,
    description_progress VARCHAR NOT NULL,
    status progress_status DEFAULT 'Open' NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);