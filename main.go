package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/user/{user}", handlers.getUserData).Methods("GET")
	http.ListenAndServe(":8080", r)
}
