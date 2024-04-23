CREATE TABLE IF NOT EXISTS customer (
    id SERIAL PRIMARY KEY,
    nationality_id INT NOT NULL,
    cst_name VARCHAR(50) NOT NULL,
    cst_dob VARCHAR(50) NOT NULL,
    cst_phoneNum VARCHAR(20) NOT NULL,
    cst_email VARCHAR(50) NOT NULL,
    FOREIGN KEY(nationality_id) REFERENCES nationality(id)
);