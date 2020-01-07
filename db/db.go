package db

import (
	"context"
	"fmt"
	"log"

	//"gopkg.in/mgo.v2"
	
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Database
)

func init() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/rest")

	// Connect to MongoDB
	if clt, err := mongo.Connect(context.TODO(), clientOptions); err != nil {
		log.Fatalln(err)
	} else {
		if err := clt.Ping(context.TODO(), nil); err != nil {
			log.Fatalln(err)
		}
		client = clt.Database("rest")

	}
	//collection := client.Collection("users")

	fmt.Println("Connected to MongoDB!")
}

// GetConnection returns a connection to the DB
func GetConnection() *mongo.Database {
	return client
}
