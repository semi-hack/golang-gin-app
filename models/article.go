package models

const (
	// Articlecollection ....
	Articlecollection = "articles"
)

// Article struct
type Article struct {
	Name string `json:"name" binding:"required"`
	Content string `json:"content" binding:"required"`
	User 
}