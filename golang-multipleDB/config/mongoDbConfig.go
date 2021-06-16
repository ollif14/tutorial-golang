package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var MongoClient *mongo.Client
var MongoCollection *mongo.Collection
var MongoCtx context.Context

const database = "employeesdatabase"
const collection = "employees"

func ConnectToMongo(){
	fmt.Println("Connecting to MongoDB...")
	MongoCtx = context.Background()
	db, err := mongo.Connect(MongoCtx, options.Client().ApplyURI("mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass&ssl=false"))
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(MongoCtx, nil); err != nil {
		log.Fatalf("Could not connect to MongoDB: %v\n", err)
	} else {
		fmt.Println("Connected to Mongodb")
	}
	MongoCollection = db.Database(database).Collection(collection)
}