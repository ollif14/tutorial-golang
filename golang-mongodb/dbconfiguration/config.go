package dbconfiguration

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

func ConnectToMongo()  *mgo.Session {
	session, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Println("session err:", err)
	}

	fmt.Println("connect")

	return session
}
