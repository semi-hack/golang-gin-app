package controller

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"auth/db"
	"auth/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"github.com/gin-gonic/gin"
)

// Createarticle ....
func Createarticle(c *gin.Context) {
	var data models.Article

	db.GetConnection().Collection(models.Articlecollection).InsertOne(context.TODO(), data)

	c.JSON(200, gin.H{"message": "article created"})

}

// Getarticle ....
func Getarticle(c *gin.Context) {
	var data models.Article

	id := struct {
		ID string `form:"_id" binding:"required"`
	}{}
	if err := c.ShouldBindQuery(&id); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	objID, err := primitive.ObjectIDFromHex(id.ID)

	if err != nil {
		c.JSON(400, gin.H{"Invalid":"Invalid ID"})
	}

	query := bson.M{"_id": objID}

	connection := db.GetConnection().Collection(models.Articlecollection)
	err = connection.FindOne(context.TODO(), query).Decode(&data)

	if err != nil {
		c.JSON(400, gin.H{"message":"error bro"})
	}
	c.JSON(200, gin.H{"Article": data})

}
