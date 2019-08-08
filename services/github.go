package services

import (
	"challenge-go-react/models"
	"encoding/json"
	"log"

	"github.com/imroc/req"
)

func GetUserRepos(username string) []models.Repo {
	log.Print("here")
	resp, _ := req.Get("https://api.github.com/users/" + username + "/repos")
	var bin, _ = resp.ToBytes()
	var repoList []models.Repo
	json.Unmarshal(bin, &repoList)
	log.Print(repoList)
	return repoList
}
