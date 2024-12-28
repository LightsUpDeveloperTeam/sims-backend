CREATE TABLE bank_account (
    id SERIAL PRIMARY KEY,
    family_id INTEGER REFERENCES family(id),
    account_name VARCHAR,
    bank_name VARCHAR
);