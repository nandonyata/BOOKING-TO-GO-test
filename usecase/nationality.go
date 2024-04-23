package usecase

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nandonyata/BOOKING-TO-GO-test/entity"
	"github.com/nandonyata/BOOKING-TO-GO-test/repository"
)

type NationalityService struct {
	Database *sql.DB
}

type Response struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

func (s *NationalityService) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newNationality entity.Nationality
	response := Response{}

	err := json.NewDecoder(r.Body).Decode(&newNationality)
	if err != nil {
		response.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	if newNationality.Nationality_code == "" || newNationality.Nationality_name == "" {
		response.Error = "Fill All Field"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	repo := repository.NationalityRepository{Database: s.Database}

	// Insert data into database
	_, err = repo.Create(&newNationality)
	if err != nil {
		response.Error = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Data = struct {
		Message string `json:"message"`
	}{Message: "Success"}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (s *NationalityService) FindAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := Response{}

	repo := repository.NationalityRepository{Database: s.Database}

	result, err := repo.FindAll()
	if err != nil {
		response.Error = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Data = result
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (s *NationalityService) FindOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	paramCode := mux.Vars(r)
	response := Response{}

	repo := repository.NationalityRepository{Database: s.Database}

	result, err := repo.FindOne(paramCode["code"])
	if err != nil {
		response.Error = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Data = result
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
