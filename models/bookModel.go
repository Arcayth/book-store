package models

type Book struct {
	ID     string `json:"id" bson:"_id,omitempty"`
	Title  string `json:"title" bson:"title,omitempty"`
	Author string `json:"author" bson:"author,omitempty"`
}
