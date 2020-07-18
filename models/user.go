package models

import(
	// "gin/db"
	// "go.mongodb.org/mongo-driver/bson"
	// "context"
	// "fmt"

)
const (
	// Usercollection ...
	Usercollection = "users"
)

// User struct ...
type User struct {
	Name string `json:"name" omitempty" binding:"required"`
	Email string `json:"email"`
	Password string `json:"password"`
	Age int `json:"age"`
	Username string `json:"username"`
}
