package model

type Employees struct {
	Id            int    `bson:"e_id" json:"id" gorm:"primary_key"`
	Email_address string `json:"email_address" bson:"email_address"`
	First_name    string `json:"first_name" bson:"first_name"`
	Last_name     string `json:"last_name" bson:"last_name"`
}
