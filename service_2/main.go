package main

import (
	"fmt"

	"github.com/joaopedro5g/hello-kafka/src/app/product"
	infra "github.com/joaopedro5g/hello-kafka/src/infra"
	"github.com/segmentio/kafka-go"
)

func main() {
	fmt.Println("Serviço rodando...")
	msgChan := make(chan *kafka.Message)
	conn := infra.NewKafkaConsumer(msgChan)

	go conn.Consumer()

	for msg := range msgChan {
		p := product.InitProduct(msg)
		p.Buy()
	}
}
