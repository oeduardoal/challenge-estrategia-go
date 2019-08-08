package models

// Repo represent repo
type Repo struct {
	ID       string   `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string   `json:"name"`
	FullName string   `json:"full_name"`
	NodeID   string   `json:"node_id"`
	Desc     string   `json:"description"`
	URL      string   `json:"url"`
	Language string   `json:"language"`
	Tags     []string `json:"tags"`
}
