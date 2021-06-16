package model

type Employee struct {
	EId           int    `bson:"e_id" json:"id" csv:"id"`
	Email_address string `json:"email_address" bson:"email_address" csv:"email_address"`
	First_name    string `json:"first_name" bson:"first_name" csv:"first_name"`
	Last_name     string `json:"last_name" bson:"last_name" csv:"last_name"`
}
