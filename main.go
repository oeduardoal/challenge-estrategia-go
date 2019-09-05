package main

import (
	"challenge-estrategia-go/handlers"
	"challenge-estrategia-go/storage"
	"net/http"

	"github.com/gorilla/mux"
)

var db *storage.DB

func init() {
	db.Connect()
}

func main() {
	r := mux.NewRouter()
	// swagger:route POST /repo repos users createRepoReq
	// Creates a new repository for the currently authenticated user.
	// If repository name is "exists", error conflict (409) will be returned.
	// responses:
	//  200: repoResp
	//  400: badReq
	//  409: conflict
	//  500: internal
	r.HandleFunc("/api/v1/user/{user}", handlers.GetUserData).Methods("GET")
	r.HandleFunc("/api/v1/tag/{user}/{repo}", handlers.GetRepoTags).Methods("GET")
	r.HandleFunc("/api/v1/tag/{user}/{repo}", handlers.SaveRepoTags).Methods("POST")
	r.HandleFunc("/api/v1/repo", handlers.GetRepoByQuery).Methods("GET")
	r.HandleFunc("/api/v1/rec", handlers.GetRecToRepo).Methods("GET")
	http.ListenAndServe(":8080", r)
}
