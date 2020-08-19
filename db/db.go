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
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/rest")
	clientOptions := options.Client().ApplyURI("mongodb+srv://semi:sengoku00@cluster0.ilogk.mongodb.net/<dbname>?retryWrites=true&w=majority")


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

type Collection struct {
	// collection is a connection to a MongoDB collection
	collection *mongo.Collection

	// database is a connection to the MongoDB database that houses the collection
	database *mongo.Database
}

// GetConnection returns a connection to the DB
func GetConnection() *mongo.Database {
	return client
}



// import "go.mongodb.org/mongo-driver/mongo"

// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// defer cancel()
// client, err := mongo.Connect(ctx, options.Client().ApplyURI(
//    "mongodb+srv://semi:<password>@cluster0.ilogk.mongodb.net/<dbname>?retryWrites=true&w=majority",
// ))
// if err != nil { log.Fatal(err) }
