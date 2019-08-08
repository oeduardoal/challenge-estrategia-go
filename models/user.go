package models

// User represent user
type User struct {
	ID       interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	Username string      `json:"username,omitempty" bson:"username,omitempty"`
	Repos    []Repo      `json:"repos"`
}
