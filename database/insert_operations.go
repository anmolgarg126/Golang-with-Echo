package database

import (
	"context"
	"echo_framework/config"
	"echo_framework/model"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

//BSON = Binary Encoded Json. Includes additional types e.g. int, long, date, floating point,
//bson.D = ordered document bson.D{{"hello", "world"}, {"foo", "bar"}}
//bson.M = unordered document/map bson.M{"hello": "world", "foo": "bar"}
//bson.A = array bson.A{"Eric", "Kyle"}
//bson.E = usually used as an element inside bson.D

// db := client. Database("tronics") collection := b.Collection("products")
// res, err := collection. InsertOne (context. Background (), trimmer)
// res, err := collection. InsertOne (context. Background(), bson.D{
//{"name", "eric"},
// {"surname","cartman"}.
// {"hobbies", bson.A{"videogame", "alexa", "kfc"}},
// })

func InsertOneUsingStruct(u model.User) error {
	collection := config.MongoDb.Collection("users")
	res, err := collection.InsertOne(context.Background(), u)
	if err != nil {
		fmt.Println("error while saving :: ", err)
		return err
	}
	fmt.Println(res)
	return nil
}

var userA = model.User{
	Name: "UserA",
	ID:   "userA",
	Jobs: []string{"police", "fireman"},
}
var userB = model.User{
	Name: "UserB",
	ID:   "userB",
	Jobs: []string{"police", "fireman"},
}

// Using bson.D

func InsertOneUsingBsonD() {
	collection := config.MongoDb.Collection("users")
	res, err := collection.InsertOne(context.Background(), bson.D{
		{"name", "bson.D_User"},
		{"_id", "bson.D_User_id"},
		{"jobs", bson.A{"abc", "def"}},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

// Using bson.M

func InsertOneUsingBsonM() {
	collection := config.MongoDb.Collection("users")
	res, err := collection.InsertOne(context.Background(), bson.M{
		"name": "bson.M_User",
		"_id":  "bson.M_User_id",
		"jobs": bson.A{"abc", "def"},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

// insert many documents

func InsertManyDocuments() {
	collection := config.MongoDb.Collection("users")
	res, err := collection.InsertMany(context.Background(), []interface{}{userA, userB})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
