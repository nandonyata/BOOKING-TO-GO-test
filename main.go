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

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World"))
	})
	r.HandleFunc("/nationality", nationalityService.Create).Methods(http.MethodPost)

	log.Printf("Listening on port %v\n", PORT)
	err := http.ListenAndServe(PORT, r)
	if err != nil {
		log.Fatalf("Error listenging on port %v, err: %v \n", PORT, err)
	}

}
