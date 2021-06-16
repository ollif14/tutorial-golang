package model

import "gopkg.in/mgo.v2/bson"

type Employee struct {
	Id            bson.ObjectId `bson:"_id"`
	E_Id int64 `bson:"e_id"`
	Email_address string `bson:"email_address"`
	First_name    string `bson:"first_name"`
	Last_name     string `bson:"last_name"`
}
