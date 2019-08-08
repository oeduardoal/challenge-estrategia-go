package services

import (
	"challenge-go-react/models"
	"encoding/json"

	"github.com/imroc/req"
)

// GetRecWords response
func GetRecWords(q string) []models.RecRequest {
	resp, _ := req.Get("https://api.datamuse.com/words?ml=" + q)
	var bin, _ = resp.ToBytes()
	var rec models.RecListRequest
	json.Unmarshal(bin, &rec.Rec)
	return rec.Rec
}
