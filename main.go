package main

import (
	"challenge-go-react/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/user/{user}", controllers.GetUserData).Methods("GET")
	http.ListenAndServe(":8080", r)
}
