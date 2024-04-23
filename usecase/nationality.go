package usecase

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nandonyata/BOOKING-TO-GO-test/entity"
	"github.com/nandonyata/BOOKING-TO-GO-test/repository"
)

type NationalityService struct {
	Database *sql.DB
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func (s *NationalityService) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := Response{}
	defer json.NewEncoder(w).Encode(res)

	var newNationality entity.Nationality

	err := json.NewDecoder(r.Body).Decode(&newNationality)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = err.Error()
		return
	}

	if newNationality.Nationality_code == "" || newNationality.Nationality_name == "" {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = "Fill All Field"
		return
	}

	repo := repository.NationalityRepository{Database: s.Database}

	// insert data to database
	result, err := repo.Create(&newNationality)
	if err != nil {
		fmt.Println("Error: ", err)

		w.WriteHeader(http.StatusBadRequest)
		res.Error = err.Error()
		return
	}

	res.Data = result
	w.WriteHeader(http.StatusCreated)
}
