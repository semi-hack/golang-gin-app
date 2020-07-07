package controller

import (
	"gin/db"
	"gin/services"
	"gin/models"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

//var user []user

// Home .....
func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"Message": "Welcome",
	})
}

// Createuser .....
func Createuser(c *gin.Context) {
	var data models.User
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"message": "Missing Field", "error": err.Error()})
		return
	}

	id, err := services.CreateUser(&data)
	if err != nil {
		c.JSON(500, gin.H{"msg": err })
	}
	c.JSON(200, gin.H{"id": id })

}

// Getuserbyid .....
func Getuserbyid(c *gin.Context) {

	var users models.User
	id := struct {
		ID string `form:"_id" binding:"required"`
	}{}
	if err := c.ShouldBindQuery(&id); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	objID, err := primitive.ObjectIDFromHex(id.ID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	query := bson.M{"_id": objID}

	connection := db.GetConnection().Collection(models.Usercollection)
	err = connection.FindOne(context.TODO(), query).Decode(&users)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"user": users})

}

// Updateuser .....
func Updateuser(c *gin.Context) {
	var users models.User

	id := struct {
		ID       string `form:"_id" binding:"required"`
		Name     string `form:"Name"`
		Password string `form:"password"`
		Age      int    `form:"age"`
		Username string `form:"username"`
	}{}

	if err := c.ShouldBindQuery(&id); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	objID, err := primitive.ObjectIDFromHex(id.ID)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	query := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"name": id.Name}}

	res := db.GetConnection().Collection(models.Usercollection).FindOneAndUpdate(context.TODO(), query, update)
	res.Decode(&users)

	c.JSON(200, gin.H{"updated": users})

}

// Deleteuser .....
func Deleteuser(c *gin.Context) {
	var users models.User
	id := struct {
		ID string `form:"_id" binding:"required"`
	}{}

	if err := c.ShouldBindQuery(&id); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	objID, err := primitive.ObjectIDFromHex(id.ID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	query := bson.M{"_id": objID}

	connection := db.GetConnection().Collection(models.Usercollection)
	connection.FindOneAndDelete(context.TODO(), query)

	c.JSON(200, gin.H{"message": users})

}
