package models

const (
	//Commentcollection ...
	Commentcollection = "comments"
)

//Comments struct ...
type Comments struct {
  Movieid string `json:"movieid"`
  User string `json:"user"`
  Comment string `json:"comment"`
  Date string `json:"date"`
}