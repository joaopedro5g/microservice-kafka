package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	MsgChan chan *kafka.Message
}

func NewKafkaConsumer(msgChan chan *kafka.Message) *KafkaConsumer {
	return &KafkaConsumer{
		MsgChan: msgChan,
	}
}

func (k *KafkaConsumer) Consumer() {
	conf := kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "test",
		GroupID: "g1",
	}

	reader := *kafka.NewReader(conf)

	for {
		m, err := reader.ReadMessage(context.Background())

		if err == nil {
			k.MsgChan <- &m
		} else {
			fmt.Print(err.Error())
		}
	}
}
