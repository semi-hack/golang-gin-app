package services

import (
	"gin/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/mongo/options"
	"gin/models"
	"context"
	"log"
	"errors"
)

// Create ...
func Create(todo *models.Todo) (primitive.ObjectID, error) {
	todo.ID = primitive.NewObjectID()
	result, err := db.GetConnection().Collection(models.Todocollection).InsertOne(context.TODO(), todo)
	if err != nil {
		log.Printf("could not create todo: %v", err)
		return primitive.NilObjectID, err
	}
	oid := result.InsertedID.(primitive.ObjectID)
	return oid, nil
}

// GetAll ...
func GetAll() ([]*models.Todo, error) {
	var todos []*models.Todo
	result, err := db.GetConnection().Collection(models.Todocollection).Find(context.TODO(), bson.M{})
	if err != nil {
		log.Printf("could not find todos", err)
	}

	err = result.All(context.TODO(), &todos)
	if err != nil {
		log.Fatal("error in decoding doc", err)
		return nil, err
	}
	return todos, nil
}

// GetTodoById ...
func GetTodoByID(id string) (*models.Todo, error) {
	var todo *models.Todo

	objID, err := primitive.ObjectIDFromHex(id)
	
	query := bson.M{"_id": objID}
	
	result := db.GetConnection().Collection(models.Todocollection).FindOne(context.TODO(), query)
	if result == nil {
		return nil, errors.New("could not find todo")
	}
	err = result.Decode(&todo)

	if err != nil {
		log.Printf("failure", err)
		return nil, err
	}
	log.Printf("Todos: %v", todo)
	return todo, nil
}

// func UpdateTodo(todo *models.Todo) (*models.Todo, error) {
// 	var updatedTodo *models.Todo

// 	//objID, _:= primitive.ObjectIDFromHex(todo.ID)

// 	query := bson.M{"_id": data}
// 	update := bson.M{
// 		"$set": todo,
// 	}
// 	upsert := true
// 	after := options.After
// 	opt := options.FindOneAndUpdateOptions {
// 		Upsert: &upsert,
// 		ReturnDocument: &after,
// 	}
	
// 	err := db.GetConnection().Collection(models.Todocollection).FindOneAndUpdate(context.TODO(), query, update, &opt).Decode(&updatedTodo)
//     if err != nil {
// 		log.Printf("could not save Todo: %v", err)
// 		return nil, err
// 	}
// 	return updatedTodo, nil
// }

// func Delete(todo *models.Todo) {
// 	result, err := db.GetConnection().Collection(models.Todocollection).FindOneAndDelete(context.TODO(), todo)
// 	if err != nil {
// 		log.Printf("could not find todo")
// 		return primitive.NilObjectID, err
// 	}
// 	return
// }