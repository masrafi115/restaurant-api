package databaseConn

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DatabaseInt() *mongo.Client {
	
	
	MongoDbAddr := "mongodb://localhost:27017"

	MongoClient, err := mongo.NewClient(options.Client().ApplyURI(MongoDbAddr))
	
	if err != nil {
		log.Fatal(err)
		fmt.Print(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = MongoClient.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nconnected to mongodb")
	return MongoClient
}

var Client *mongo.Client = DatabaseInt()

func OpenDBCol(client *mongo.Client, collName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurent").Collection(collName)

	return collection
}

func CloseColl(client *mongo.Client) {
	client.Disconnect(context.TODO())
}