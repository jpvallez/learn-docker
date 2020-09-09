package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/jpvallez/learn-docker/datasources/apifootball"
)

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

	fmt.Println("Connected to MongoDB!!!")

	collection := client.Database("test").Collection("players")

	var premierLeagueTeams []apifootball.Team

	premierLeagueTeams, err = apifootball.GetTeamsFromLeagueId("2", false)
	if err != nil {
		log.Fatal(err)
	}

	// Let's try and stuff the teams one by one into mongodb

	for _, team := range premierLeagueTeams {
		insertResult, err := collection.InsertOne(context.TODO(), team)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted a team, insert ID: ", insertResult.InsertedID)
	}

	// Now for testing sake let's get a team back from the database and print it.

	filter := bson.M{"name": "Southampton"}

	var result apifootball.Team

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)

}
