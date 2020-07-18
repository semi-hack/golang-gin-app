package controller

import (
	//"context"
	//"context"
	//"fmt"
	//"gin/db"
	//"gin/models"
	"gin/services"

	"github.com/gin-gonic/gin"
	//"go.mongodb.org/mongo-driver/bson"
)

//var userModel = new(models.UserModel)

// Login controller
func Login(c *gin.Context) {
	//var user *models.User
	data := struct {
		username string `form:"username" binding:"required"`
		password string `form:"password" binding:"required"`
		//ID string `form:"_id" binding:"required"`

	}{}

	if err := c.ShouldBindQuery(&data); err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	zh, err := services.GetUser(data.username)
	if err != nil {
		c.JSON(400, gin.H{"message": "error finding user"})
		return
	}
	//fmt.Println(user)


	hashedPassword := []byte(zh.Password)
	password := []byte(data.password)

	err = services.PasswordCompare(hashedPassword, password)
	if err != nil {
		c.JSON(403, gin.H{"message": "Invalid user credentials"})
		return
	}

	toks, err := services.GenerateToken(data.username)
	if err != nil {
		c.JSON(403, gin.H{"message": "There was a problem logging you in, try again later"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "success", "token": toks})
}
