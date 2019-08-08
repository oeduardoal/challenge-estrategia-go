package main

import (
	"challenge-go-react/handlers"
	"challenge-go-react/storage"
	"net/http"

	"github.com/gorilla/mux"
)

var db *storage.DB

func init() {
	db.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/user/{user}", handlers.GetUserData).Methods("GET")
	r.HandleFunc("/api/v1/tag/{user}/{repo}", handlers.GetRepoTags).Methods("GET")
	r.HandleFunc("/api/v1/tag/{user}/{repo}", handlers.SaveRepoTags).Methods("POST")
	r.HandleFunc("/api/v1/repo", handlers.GetRepoByQuery).Methods("GET")
	r.HandleFunc("/api/v1/rec", handlers.GetRecToRepo).Methods("GET")
	http.ListenAndServe(":8080", r)
}
