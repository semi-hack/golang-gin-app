package controller

import (
	"gin/services"
	"gin/db"
	"gin/models"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

)

// Login controller
func Login(c *gin.Context) {
	var user models.User
	data := models.User{}
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()} )
		return
	}

	connection := db.GetConnection().Collection(models.Usercollection)
	err = connection.FindOne(context.TODO(), bson.M{data.Email: user.Email}).Decode(&user)

	if err != nil {
		c.JSON(401, gin.H{"message": "error finding user"})
		return
	}

	hashedPassword := []byte(user.Password)
	Password := []byte(data.Password)

	err = services.PasswordCompare(hashedPassword, Password)
	if err != nil {
		c.JSON(403, gin.H{"message": "Invalid user credentials"})
		return
	}

	toks, err := services.GenerateToken(data.Email)
	if err != nil {
        c.JSON(403, gin.H{"message": "There was a problem logging you in, try again later"})
        c.Abort()
        return
    }

    c.JSON(200, gin.H{"message": "Log in success", "token": toks})
}