package consumer

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"golang-kafka-2/config"
	"golang-kafka-2/model"
)

func NewConsumer() (sarama.Consumer, error){
	var brokers = []string{"localhost:9092"}
	var consumer, err = sarama.NewConsumer(brokers, nil)

	return consumer, err
}

func Consume(topic string, consumer sarama.Consumer) {
	partitionList, err := consumer.Partitions(topic) //get all partitions on the given topic
	if err != nil {
		fmt.Println("Error retrieving partitionList ", err)
	}
	initialOffset := sarama.OffsetOldest //get offset for the oldest message on the topic

	for _, partition := range partitionList {
		pc, _ := consumer.ConsumePartition(topic, partition, initialOffset)
		for message := range pc.Messages(){
			receiveMsg(message)
		}
	}
}

func receiveMsg(message *sarama.ConsumerMessage) {
	var save model.Employees
	json.Unmarshal(message.Value, &save)
	fmt.Printf("%d : %s : %s : %s\n", save.Id, save.Email_address, save.First_name, save.Last_name)

	_, err := config.Employeedb.InsertOne(config.MongoCtx, save)
	if err != nil {
		panic(err)
	}
}
