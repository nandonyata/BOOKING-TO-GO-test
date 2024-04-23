package repository

import (
	"database/sql"

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

func (r *CustomerRepository) FindAll() (interface{}, error) {
	query := `
		SELECT
			c.id,
			c.nationality_id,
			c.cst_name,
			c.cst_dob,
			c.cst_phoneNum,
			c.cst_email,
			f.fl_relation,
			f.fl_name,
			f.fl_dob,
			n.nationality_name
		FROM
			customer c
		JOIN
			family_list f ON c.id = f.cst_id
		JOIN
			nationality n ON c.nationality_id = n.id
		ORDER BY
			c.id DESC
	`

	rows, err := r.Database.Query(query)
	if err != nil {
		return nil, err
	}

	var response []model.ResponseFetchCustomer
	var currentCustId int
	var currentCust model.ResponseFetchCustomer

	for rows.Next() {

		var customer entity.Customer
		var nationality string
		var family_list model.FamilyList

		err := rows.Scan(&customer.Id, &customer.Nationality_id, &customer.Cst_name, &customer.Cst_dob, &customer.Cst_phoneNum, &customer.Cst_email, &family_list.Relation, &family_list.Name, &family_list.Dob, &nationality)

		if err != nil {
			return nil, err
		}

		if customer.Id != currentCustId {
			if currentCustId != 0 {
				response = append(response, currentCust)
			}

			currentCustId = customer.Id
			currentCust = model.ResponseFetchCustomer{
				Id:          customer.Id,
				Nationality: nationality,
				Name:        customer.Cst_name,
				Dob:         customer.Cst_dob,
				Phone:       customer.Cst_phoneNum,
				Email:       customer.Cst_email,
				Family_list: []model.FamilyList{family_list},
			}
		} else {
			currentCust.Family_list = append(currentCust.Family_list, family_list)
		}

	}

	if currentCustId != 0 {
		response = append(response, currentCust)
	}

	return response, nil

}

func (r *CustomerRepository) FindOne(id int) (interface{}, error) {
	query := `
		SELECT
			c.id,
			c.nationality_id,
			c.cst_name,
			c.cst_dob,
			c.cst_phoneNum,
			c.cst_email,
			f.fl_relation,
			f.fl_name,
			f.fl_dob,
			n.nationality_name
		FROM
			customer c
		JOIN
			family_list f ON c.id = f.cst_id
		JOIN
			nationality n ON c.nationality_id = n.id
		WHERE
			c.id = $1
		ORDER BY
			c.id DESC
	`

	rows, err := r.Database.Query(query, id)
	if err != nil {
		return nil, err
	}

	var response model.ResponseFetchCustomer
	var currentCustId int

	for rows.Next() {

		var customer entity.Customer
		var nationality string
		var family_list model.FamilyList

		err := rows.Scan(&customer.Id, &customer.Nationality_id, &customer.Cst_name, &customer.Cst_dob, &customer.Cst_phoneNum, &customer.Cst_email, &family_list.Relation, &family_list.Name, &family_list.Dob, &nationality)

		if err != nil {
			return nil, err
		}

		if currentCustId == 0 {
			currentCustId = customer.Id
			response = model.ResponseFetchCustomer{
				Id:          customer.Id,
				Nationality: nationality,
				Name:        customer.Cst_name,
				Dob:         customer.Cst_dob,
				Phone:       customer.Cst_phoneNum,
				Email:       customer.Cst_email,
				Family_list: []model.FamilyList{family_list},
			}
		} else {
			response.Family_list = append(response.Family_list, family_list)
		}
	}

	return response, nil

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
