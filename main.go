package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nandonyata/BOOKING-TO-GO-test/config"
)

func main() {
	config.ConnectDB()

	r := mux.NewRouter()
	// r.HandleFunc("/", HomeHandler)
	// r.HandleFunc("/products", ProductsHandler)
	// r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)

}
