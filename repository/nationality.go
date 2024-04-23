package repository

import (
	"database/sql"
	"errors"

	"github.com/nandonyata/BOOKING-TO-GO-test/entity"
)

type NationalityRepository struct {
	Database *sql.DB
}

func (r *NationalityRepository) Create(in *entity.Nationality) (interface{}, error) {
	query := `
		INSERT INTO nationality (nationality_name, nationality_code)
		VALUES ($1, $2)
	`

	result, err := r.Database.Exec(query, in.Nationality_name, in.Nationality_code)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *NationalityRepository) FindAll() (interface{}, error) {
	query := `
		SELECT nationality_name, nationality_code FROM nationality
		ORDER BY nationality_name ASC
	`

	rows, err := r.Database.Query(query)
	if err != nil {
		return nil, err
	}

	var data []entity.Nationality

	for rows.Next() {
		nationality := entity.Nationality{}

		if err = rows.Scan(&nationality.Nationality_name, &nationality.Nationality_code); err != nil {
			return nil, err
		}

		data = append(data, nationality)
	}

	return data, nil
}

func (r *NationalityRepository) FindById(in int) (interface{}, error) {
	query := `
		SELECT nationality_name, nationality_code FROM nationality
		WHERE id = $1
	`

	data := entity.Nationality{}

	err := r.Database.QueryRow(query, in).Scan(&data.Nationality_name, &data.Nationality_code)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("data not found")
		}
	}

	return data, nil
}
