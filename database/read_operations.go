package database

import (
	"context"
	"echo_framework/config"
	"echo_framework/model"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

// Equality operator using FindOne

func FindOneByName(name string) (model.User, error) {
	var u model.User
	collection := config.MongoDb.Collection("users")

	// fmt.Printf("%T", collection)
	if err := collection.FindOne(context.Background(), bson.M{"name": name}).Decode(&u); err != nil {
		fmt.Println("error while reading :: ", err)
		return u, err
	}
	fmt.Println(u)
	return u, nil
}

func FindOneByID(ID string) (model.User, error) {
	var u model.User
	collection := config.MongoDb.Collection("users")
	if err := collection.FindOne(context.Background(), bson.M{"_id": ID}).Decode(&u); err != nil {
		fmt.Println("error while reading :: ", err)
		return u, err
	}
	fmt.Println(u)
	return u, nil
}

// Comparision operator using Find

func FindByNumberOfJobs(num int) ([]model.User, error) {
	result := make([]model.User, 0)

	collection := config.MongoDb.Collection("users")
	// findCursor, err := collection.Find(context.Background(), bson.M{"jobs": bson.M{"$size": bson.M{"$gt": num}}})
	findCursor, err := collection.Find(context.Background(), bson.M{"jobs": bson.M{"$size": num}})
	if err != nil {
		fmt.Println("error in FindByComparisingJobs :: ", err)
		return result, err
	}

	fmt.Println("findCursor :: ", findCursor)

	return iterateToFindList(findCursor)
}

// Logical operator using Find

func FindByExperienceAndSalary(salary float64, experience int) ([]model.User, error) {
	collection := config.MongoDb.Collection("users")
	logicFilter := bson.M{
		"$and": bson.A{
			bson.M{"salary": bson.M{"$gt": salary}},
			bson.M{"experience": bson.M{"$gt": experience}},
		},
	}

	result := make([]model.User, 0)

	fmt.Println("---------- Logical operator using Find ----------")
	findLogicRes, err := collection.Find(context.Background(), logicFilter)

	// fmt.Printf("%T", findLogicRes)

	if err != nil {
		fmt.Println("error in FindByExperienceAndSalary :: ", err)
		return result, err
	}

	return iterateToFindList(findLogicRes)
}

func iterateToFindList(cursor *mongo.Cursor) ([]model.User, error) {
	defer cursor.Close(context.Background())
	result := make([]model.User, 0)

	for cursor.Next(context.Background()) {
		var findLogic model.User

		if err := cursor.Decode(&findLogic); err != nil {
			fmt.Println(err)
			return result, err
		}
		fmt.Println(findLogic)
		result = append(result, findLogic)
	}
	return result, nil
}

// Element operator using Find

func FindUsersWhoseNameExists() ([]model.User, error) {
	collection := config.MongoDb.Collection("users")

	elementFilter := bson.M{
		"name": bson.M{"$exists": true},
	}

	result := make([]model.User, 0)

	fmt.Println("------------Element operator using Find ---------")
	findElementRes, err := collection.Find(context.Background(), elementFilter)

	if err != nil {
		fmt.Println("error in FindUsersWhoseNameExists :: ", err)
		return result, err
	}
	return iterateToFindList(findElementRes)
}

// Array operator using Find

func FindUsersWorkedInGivenOrganisation(organisation string) ([]model.User, error) {
	collection := config.MongoDb.Collection("users")

	arrayFilter := bson.M{
		"jobs": bson.M{"$all": bson.A{organisation}},
	}

	result := make([]model.User, 0)

	fmt.Println("------------Array operator using Find ---------")
	findElementRes, err := collection.Find(context.Background(), arrayFilter)

	if err != nil {
		fmt.Println("error in FindUsersWorkedInGivenOrganisation :: ", err)
		return result, err
	}
	return iterateToFindList(findElementRes)
}
