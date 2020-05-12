package kafka

import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func Producer(ipAddress string) *kafka.Producer {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": ipAddress})
	if err != nil {
		panic(err)
	}
	return producer
}

func SendMessage(producer *kafka.Producer, topic string, message []byte, key string) {
	messageChan := make(chan kafka.Event)
	err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
		Key: 			[]byte(key),
		Headers:        []kafka.Header{{Key: "myTestHeader", Value: []byte("header values are binary")}},
	}, messageChan)
	if err != nil {
		panic(err)
	}

	e := <-messageChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}
	close(messageChan)
}
