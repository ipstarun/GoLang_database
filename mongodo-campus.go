package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbHost = "localhost"
	dbName = "threejoin"
)

const numGoroutines = 4 // Number of concurrent goroutines

var counter int // Global counter variable

func main() {
	//Contexts are used for controlling and managing important aspects of reliable applications, 
	//such as cancellation and data sharing in concurrent programming

	//method to specify the connection details.
	clientOptions := options.Client().ApplyURI("mongodb://" + dbHost)
	client, err := mongo.Connect(context.Background(), clientOptions)

	

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()
		
	collection := client.Database(dbName).Collection("counter_collection")

	// Create the collection and insert one initial record with a counter value of 0
	_, err = collection.InsertOne(context.Background(), bson.M{"val": 0})
	if err != nil {
		log.Fatal(err)
	}

	// Channel to receive completion signals from goroutines
	done := make(chan bool)

	fmt.Println("dataBase connected successfully !!!!!!")

	// Start multiple goroutines to increment the counter concurrently
	for i := 0; i < numGoroutines; i++ {
		fmt.Println("starting time: " , time.Now())
		go func() {
			var timeSecond = time.Second
			startTime := time.Now()
			for time.Since(startTime) < timeSecond {
				_, err = collection.UpdateOne(
					context.Background(),
					bson.M{},
					bson.M{"$inc": bson.M{"val": 1}},
				)
				if err != nil {
					log.Fatal(err)
				}
				counter++
			}
			done <- true
			fmt.Printf("Incremented counter %d times.\n",  counter)
		}()
	}

	// Wait for all goroutines to complete
	for i := 0; i < numGoroutines; i++ {
		<-done
	}

	fmt.Println("Completion Time: ", time.Now())
	fmt.Printf("Final Counter Value: %d\n", counter)
}
