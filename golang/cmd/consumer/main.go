package main

import (
	"context"
	"fmt"
	"os"

	"kafkademo/kafka_go"

	_ "github.com/HXSecurity/DongTai-agent-go/run/base"
	_ "github.com/HXSecurity/DongTai-agent-go/run/kafkaGo"
)

func main() {
	kafkaTopic := os.Getenv("KAFKA_TOPIC")
	if kafkaTopic == "" {
		kafkaTopic = "addUserV2"
	}

	// connect to kafka: x.x.x.x:9092
	kafkaBroker := os.Getenv("KAFKA_BROKER")
	r := kafka_go.ConnectConsumer(kafkaBroker, kafkaTopic)
	defer r.Close()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
