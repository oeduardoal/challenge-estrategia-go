package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	var a DB
	r.HandleFunc("/api/v1/user/{user}", controllers).Methods("GET")
	http.ListenAndServe(":8080", r)
}
