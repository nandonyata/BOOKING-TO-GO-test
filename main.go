package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nandonyata/BOOKING-TO-GO-test/config"
)

var PORT = ":3002"

func main() {
	config.ConnectDB()

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World"))
	})

	log.Printf("Listening on port %v\n", PORT)
	err := http.ListenAndServe(PORT, r)
	if err != nil {
		log.Fatalf("Error listenging on port %v, err: %v \n", PORT, err)
	}

}
