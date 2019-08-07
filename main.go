package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

var r *mux.Router

func init() {
	r.HandleFunc("/api/v1/user/{user}", handlers.getUserData).Methods("GET")
}

func main() {
	r := mux.NewRouter()
	http.ListenAndServe(":8080", r)
}
