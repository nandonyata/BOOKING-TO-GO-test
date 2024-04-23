CREATE TABLE IF NOT EXISTS family_list (
    id SERIAL PRIMARY KEY,
    cst_id INT NOT NULL,
    fl_relation VARCHAR(50) NOT NULL,
    fl_name VARCHAR(50) NOT NULL,
    fl_dob VARCHAR(50) NOT NULL,
    FOREIGN KEY(cst_id) REFERENCES customer(id)
);