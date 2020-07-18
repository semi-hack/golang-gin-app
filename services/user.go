package services

import (
	"gin/db"
	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/mongo/options"
	"gin/models"
	"context"
	"log"
)

// CreateUser ...
func CreateUser(user *models.User) (primitive.ObjectID, error){
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return primitive.NilObjectID, err
	}
	user.Password = string(hash)
	result, err := db.GetConnection().Collection(models.Usercollection).InsertOne(context.TODO(), user)
    if err != nil {
		log.Printf("could not create todo: %v", err)
		return primitive.NilObjectID, err
	}
	oid := result.InsertedID.(primitive.ObjectID)
	return oid, nil
}

// GetUserByID ...
func GetUserByID(id string) (*models.User, error) {
	var user *models.User

	objID, err := primitive.ObjectIDFromHex(id)
	
	query := bson.M{"_id": objID}

	err = db.GetConnection().Collection(models.Usercollection).FindOne(context.TODO(), query).Decode(&user)
	if err != nil {
		log.Println("failure", err)
		return nil, err
	}
	return user, nil
}
