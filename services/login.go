package services

import (
	"context"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"

	"gin/db"
	"gin/models"
)


// GetUser ...
func GetUser(foo string) (*models.User, error) {
	var user *models.User
	//objID, err := primitive.ObjectIDFromHex(ID)
	query := bson.M{"username": foo}
	err := db.GetConnection().Collection(models.Usercollection).FindOne(context.TODO(), query).Decode(&user)
	if err != nil {
		log.Println("failure", err)
		return nil, err
	} 
	fmt.Println(user)
	return user, nil

}
