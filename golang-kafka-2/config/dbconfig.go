package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Db *mongo.Client
var Employeedb *mongo.Collection
var MongoCtx context.Context

const database = "employeesdatabase"
const collection = "employees"

func ConnectToMongo() (string, error){
	MongoCtx = context.Background()
	db, err := mongo.Connect(MongoCtx, options.Client().ApplyURI("mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass&ssl=false"))
	if err != nil {
		log.Fatal(err)
	}
	var msg string
	if err := db.Ping(MongoCtx, nil); err != nil {
		msg = "Could not connect to MongoDB: %v\n"
	} else {
		msg = "Success Connected to Mongodb"
	}
	Employeedb = db.Database(database).Collection(collection)
	return msg, err
}
