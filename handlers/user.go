package handlers

import (
	"challenge-go-react/models"
	"challenge-go-react/services"
	"challenge-go-react/storage"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetUserData return the cellar
func GetUserData(w http.ResponseWriter, r *http.Request) {
	var db *storage.DB
	params := mux.Vars(r)
	user, err := db.FindUserByUsername(params["user"])
	if err != nil {
		var newUser models.User
		newUser.Username = params["user"]
		newUser.Repos = services.GetUserRepos(params["user"])
		mongoInsert := db.InsertUser(&newUser)
		newUser.ID = mongoInsert.InsertedID
		json.NewEncoder(w).Encode(newUser)
	} else {
		json.NewEncoder(w).Encode(user)
	}
}
