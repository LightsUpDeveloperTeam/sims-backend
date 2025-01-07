CREATE TABLE shift_schedules (
    id SERIAL PRIMARY KEY,
    shift_id INT NOT NULL REFERENCES shifts(id),
    day_of_week INT NOT NULL CHECK (day_of_week >= 0 AND day_of_week <= 6),
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);