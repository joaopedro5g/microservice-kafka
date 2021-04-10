package configuration

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type Kafka struct {
	Broker  string
	Topic   string
	GroupID string
}

func StartKafka(topic string, groupId string) *Kafka {
	return &Kafka{
		Broker:  "localhost:9092",
		Topic:   topic,
		GroupID: groupId,
	}
}

func (k *Kafka) KafkaCon() *kafka.Conn {
	conn, err := kafka.DialLeader(context.Background(), "tcp", k.Broker, k.Topic, 1)
	if err != nil {
		log.Fatal(err.Error())
	}

	return conn
}
