CREATE TABLE attendance_records (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    date DATE NOT NULL,
    clock_in_time TIMESTAMP,
    clock_out_time TIMESTAMP,
    clock_in_latitude DECIMAL(9,6),
    clock_in_longitude DECIMAL(9,6),
    clock_out_latitude DECIMAL(9,6),
    clock_out_longitude DECIMAL(9,6),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);