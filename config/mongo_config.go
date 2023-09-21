package config

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDb *mongo.Database

func mongoConfig() {
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", Configuration.Database.DBHost, Configuration.Database.DBPort)))
	if err != nil {
		fmt.Println(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}

	MongoDb = client.Database(Configuration.Database.DBName)
	// collection := db.Database(Configuration.Database.DBCollection)

	// res, err := collection.InsertOne(context.Background(), "")

}
