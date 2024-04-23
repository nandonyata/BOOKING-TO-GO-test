CREATE TABLE IF NOT EXISTS nationality (
    id SERIAL PRIMARY KEY,
    nationality_name VARCHAR(255) NOT NULL,
    nationality_code VARCHAR(255) NOT NULL
);