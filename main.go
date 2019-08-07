package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oeduardoal/challenge-go-react/database"
)

func main() {
	r := mux.NewRouter()
	var a database.DB
	var b controllers.eaoooooo
	r.HandleFunc("/api/v1/user/{user}", eaoooooo).Methods("GET")
	http.ListenAndServe(":8080", r)
}
