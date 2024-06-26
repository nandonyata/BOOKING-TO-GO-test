package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nandonyata/BOOKING-TO-GO-test/config"
	"github.com/nandonyata/BOOKING-TO-GO-test/usecase"
)

var PORT = ":3002"

func main() {
	config.ConnectDB()

	nationalityService := usecase.NationalityService{Database: config.DB}
	customerService := usecase.CustomerService{Database: config.DB}

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World"))
	})
	r.HandleFunc("/nationality", nationalityService.Create).Methods(http.MethodPost)
	r.HandleFunc("/nationality", nationalityService.FindAll).Methods(http.MethodGet)
	r.HandleFunc("/nationality/{code}", nationalityService.FindOne).Methods(http.MethodGet)
	r.HandleFunc("/customer", customerService.Create).Methods(http.MethodPost)
	r.HandleFunc("/customer", customerService.FindAll).Methods(http.MethodGet)
	r.HandleFunc("/customer/{id}", customerService.FindOne).Methods(http.MethodGet)
	r.HandleFunc("/customer/{id}", customerService.Update).Methods(http.MethodPut)
	r.HandleFunc("/customer/{id}", customerService.Delete).Methods(http.MethodDelete)

	log.Printf("Listening on port %v\n", PORT)
	err := http.ListenAndServe(PORT, r)
	if err != nil {
		log.Fatalf("Error listenging on port %v, err: %v \n", PORT, err)
	}

}
