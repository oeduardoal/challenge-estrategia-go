package models

// TagRequest request
type TagRequest struct {
	Reponame string   `json:"repo"`
	Tags     []string `json:"tags"`
}

// RecListRequest request
type RecListRequest struct {
	ID     interface{}  `json:"id,omitempty" bson:"_id,omitempty"`
	Search string       `json:"search" bson:"search,omitempty"`
	Rec    []RecRequest `json:"rec" bson:"rec,omitempty"`
}

// RecRequest request
type RecRequest struct {
	Word  string `json:"word" bson:"word,omitempty"`
	Score int    `json:"score" bson:"score,omitempty"`
}
