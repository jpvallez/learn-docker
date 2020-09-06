package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Player struct {
	Name        string
	SquadNumber int
	Team        Team
}

type Team string

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://db:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("test").Collection("players")

	// Let's put Danny Ings in the database
	dannyIngs := Player{"Danny Ings", 10, "Southampton"}

	insertResult, err := collection.InsertOne(context.TODO(), dannyIngs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a document: ", insertResult.InsertedID)

	// Now let's get it back from the database and print it.

	filter := bson.M{"name": "Danny Ings"}

	var result Player

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)

}
