package main

import (
	"fmt"
	"golang-kafka-2/config"
	consumer2 "golang-kafka-2/consumer"
)

func main()  {
	msg, err := config.ConnectToMongo()
	if err != nil {
		fmt.Println("Could not create consumer: ", err)
	}
	consumer, err := consumer2.NewConsumer()
	if err != nil {
		fmt.Println("Could not create consumer: ", err)
	}
	fmt.Println(msg)
	consumer2.Consume("test-topic-1", consumer)
}
