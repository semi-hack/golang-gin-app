package models

const (
	// Usercollection ...
	Usercollection = "users"
)

// User struct ...
type User struct {
	Name string `json:"name" omitempty" binding:"required"`
	Email string `json:"email"`
	Password string `json:"password" bson:"password"`
	Age int `json:"age" bson: "age"`
	Username string `json:"username" bson: "username"`
}

