package producer

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"golang-kafka-2/model"
)

var brokers = []string{"localhost:9092"}

func NewProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)

	return producer, err
}

func Produce(topic string, e model.Employees) *sarama.ProducerMessage {
	jsonString, err := json.Marshal(e)
	if err != nil{
		return nil
	}
	empStr := string(jsonString)
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(empStr),
	}
	return msg
}