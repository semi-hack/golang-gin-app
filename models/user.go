package models

const (
	// Usercollection ...
	Usercollection = "users"
)

// User struct
type User struct {
	Name string `json:"name,omitempty" binding:"required"`
	Password string `json:"password"`
	Age int `json:"age"`
	Username string `json:"username"`
}
