package services

import (
	"gin/db"
	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/mongo/options"
	"gin/models"
	"context"
	"log"
	//"errors"
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