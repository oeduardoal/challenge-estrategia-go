package controllers

import (
	"challenge-go-react/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetUserData return the cellar
func GetUserData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := services.GetUserRepos(params["user"])
	json.NewEncoder(w).Encode(name)
}
