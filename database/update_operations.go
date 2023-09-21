package database

import (
	"context"
	"echo_framework/config"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

//Update operator for Field

func UpdateName(oldName, UpdatedName string) {
	collection := config.MongoDb.Collection("users")
	fmt.Println("-- Update Operator for field --")
	updateFieldCon := bson.M{"$set": bson.M{"name": UpdatedName}}
	updateFieldRes, err := collection.UpdateMany(context.Background(), bson.M{"name": oldName}, updateFieldCon)
	if err != nil {
		fmt.Println("error while updating name :: ", err)
	}
	fmt.Println(updateFieldRes)
	fmt.Println(updateFieldRes.ModifiedCount)
}

//Update operator for Array

func UpdateJobs() {
	collection := config.MongoDb.Collection("users")
	fmt.Println("----Update Operator for Array---")
	updateArrayCon := bson.M{"$addToSet": bson.M{"accessories": "manual"}}
	updateArrayRes, err := collection.UpdateMany(context.Background(), arrayFilter, updateArrayCon)
	if err != nil {
		fmt.Println("error while updating jobs :: ", err)
	}
	fmt.Println(updateArrayRes)
	fmt.Println(updateArrayRes.ModifiedCount)
}

//Update operator for field multiple operators

func UpdateSalary() {
	collection := config.MongoDb.Collection("users")
	fmt.Println("----Update Operator for field multiple operators----")
	incCon := bson.M{
		"$mul": bson.M{
			"price": 1.20,
		},
	}
	incRes, err := collection.UpdateMany(context.Background(), bson.M{}, incCon)
	if err != nil {
		fmt.Println("error while updating salary :: ", err)
	}
	fmt.Println(incRes)
	fmt.Println(incRes.MatchedCount)
}

// Delete operation

func DeleteById(id string) error {
	collection := config.MongoDb.Collection("users")
	fmt.Println("----Delete operation----")

	delRes, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		fmt.Println("error while deleting by id :: ", err)
		return err
	}
	fmt.Println(delRes)
	fmt.Println(delRes.DeletedCount)
	if delRes.DeletedCount == 0 {
		return errors.New("no record exists for given id")
	}
	return nil
}

var arrayFilter = bson.M{
	"jobs": bson.M{"$all": bson.A{"name"}},
}
