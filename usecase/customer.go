package usecase

import (
	"database/sql"
	"net/http"
)

type CustomerService struct {
	Database *sql.DB
}

func (s *CustomerService) Create(w http.ResponseWriter, r *http.Request) {
}

func (s *CustomerService) FindAll(w http.ResponseWriter, r *http.Request) {
}

func (s *CustomerService) FindOne(w http.ResponseWriter, r *http.Request) {
}

func (s *CustomerService) Update(w http.ResponseWriter, r *http.Request) {
}

func (s *CustomerService) Delete(w http.ResponseWriter, r *http.Request) {
}
