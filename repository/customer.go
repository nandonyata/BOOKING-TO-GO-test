package repository

import (
	"database/sql"
	"errors"

	"github.com/nandonyata/BOOKING-TO-GO-test/entity"
	"github.com/nandonyata/BOOKING-TO-GO-test/model"
)

type CustomerRepository struct {
	Database *sql.DB
}

func (r *CustomerRepository) Create(in *model.ReqBodyCreateCustomer) (interface{}, error) {
	query := `
		INSERT INTO customer (nationality_id, cst_name, cst_dob, cst_phoneNum, cst_email)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	var custId int
	err := r.Database.QueryRow(query, in.Nationality_id, in.Name, in.Dob, in.Phone, in.Email).Scan(&custId)
	if err != nil {
		return nil, err
	}

	for _, v := range in.Family_list {
		query = `
			INSERT INTO family_list (cst_id, fl_relation, fl_name, fl_dob)
			VALUES ($1, $2, $3, $4)
		`

		_, err = r.Database.Exec(query, custId, v.Relation, v.Name, v.Dob)
		if err != nil {
			return nil, err
		}
	}

	return "Success", nil
}

func (r *CustomerRepository) FindAll() ([]entity.Customer, error) {
	query := `
		SELECT * FROM customer
		ORDER BY id DESC
	`

	rows, err := r.Database.Query(query)
	if err != nil {
		return nil, err
	}

	var data []entity.Customer
	for rows.Next() {
		cust := entity.Customer{}

		if err = rows.Scan(&cust.Id, &cust.Nationality_id, &cust.Cst_name, &cust.Cst_dob, &cust.Cst_phoneNum, &cust.Cst_email); err != nil {
			return nil, err
		}

		data = append(data, cust)
	}

	return data, nil
}

func (r *CustomerRepository) FindOne(id int) (interface{}, error) {
	query := `
		SELECT * FROM customer
		WHERE id = $1
	`

	cust := entity.Customer{}

	err := r.Database.QueryRow(query, id).Scan(&cust.Id, &cust.Nationality_id, &cust.Cst_name, &cust.Cst_dob, &cust.Cst_phoneNum, &cust.Cst_email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("data not found")
		}
	}

	return cust, nil

}

func (r *CustomerRepository) Update(id int, in entity.Customer) (interface{}, error) {
	query := `
		UPDATE customer
		SET	nationality_id = $2, cst_name = $3, cst_dob = $4, cst_phoneNum= $5, cst_email = $6
		WHERE id = $1
	`

	resp, err := r.Database.Exec(query, id, in.Nationality_id, in.Cst_name, in.Cst_dob, in.Cst_phoneNum, in.Cst_email)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *CustomerRepository) Delete(id int) (interface{}, error) {
	query := `
		DELETE FROM customer
		WHERE id = $1
	`

	resp, err := r.Database.Exec(query, id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
