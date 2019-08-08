package handlers

import (
	"challenge-go-react/models"
	"challenge-go-react/services"
	"challenge-go-react/storage"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetRepoTags return the cellar
func GetRepoTags(w http.ResponseWriter, r *http.Request) {
	var db *storage.DB
	params := mux.Vars(r)
	user, err := db.FindTagsByUsernameAndReponame(params["user"], params["repo"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(user)
	}
}

func removeDuplicate(elements []string) []string {
	encountered := map[string]bool{}
	for v := range elements {
		encountered[elements[v]] = true
	}
	result := []string{}
	for key, _ := range encountered {
		result = append(result, key)
	}
	return result
}

// SaveRepoTags return the cellar
func SaveRepoTags(w http.ResponseWriter, r *http.Request) {
	var db *storage.DB
	params := mux.Vars(r)
	var t models.TagRequest
	json.NewDecoder(r.Body).Decode(&t)
	t.Tags = removeDuplicate(t.Tags)
	user, err := db.FindTagsAndUpdate(params["user"], params["repo"], t)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(user)
	}
}

// GetRecToRepo return the cellar
func GetRecToRepo(w http.ResponseWriter, r *http.Request) {
	var db *storage.DB
	q := r.FormValue("q")
	rec, err := db.FindRecByQuery(q)
	if err != nil {
		var t models.RecListRequest
		t.Search = q
		t.Rec = services.GetRecWords(q)
		if len(t.Rec) > 0 {
			db.InsertRec(&t)
		}
		json.NewEncoder(w).Encode(t)
	} else {
		json.NewEncoder(w).Encode(rec)
	}
}
