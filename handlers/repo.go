package handlers

import (
	"challenge-go-react/storage"
	"encoding/json"
	"net/http"
)

// GetRepoByQuery return the cellar
func GetRepoByQuery(w http.ResponseWriter, r *http.Request) {
	var db *storage.DB
	q := r.FormValue("q")
	user := r.FormValue("user")
	repos, err := db.FindRepoByQuery(user, q)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(repos)
	}
}
