package usecase

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/nandonyata/BOOKING-TO-GO-test/model"
	"github.com/nandonyata/BOOKING-TO-GO-test/repository"
)

type CustomerService struct {
	Database *sql.DB
}

func (s *CustomerService) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var customerBody model.ReqBodyCreateCustomer
	response := Response{}

	err := json.NewDecoder(r.Body).Decode(&customerBody)
	if err != nil {
		response.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	if customerBody.Name == "" || customerBody.Dob == "" || customerBody.Nationality_id == 0 {
		response.Error = "Fill All Field"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	repo := repository.CustomerRepository{Database: s.Database}
	_, err = repo.Create(&customerBody)
	if err != nil {
		response.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Data = struct {
		Message string `json:"message"`
	}{Message: "Success"}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (s *CustomerService) FindAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := Response{}

	repo := repository.CustomerRepository{Database: s.Database}

	result, err := repo.FindAll()
	if err != nil {
		response.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Data = result
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (s *CustomerService) FindOne(w http.ResponseWriter, r *http.Request) {
}

func (s *CustomerService) Update(w http.ResponseWriter, r *http.Request) {
}

func (s *CustomerService) Delete(w http.ResponseWriter, r *http.Request) {
}
