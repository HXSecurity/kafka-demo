package kafka_go

import (
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

func ConnectProducer(kafkaBrokerUrl string) *kafka.Writer {
	w := &kafka.Writer{
		Addr: kafka.TCP(kafkaBrokerUrl),
		// NOTE: When Topic is not defined here, each Message must define it instead.
		Balancer: &kafka.LeastBytes{},
	}

	return w
}

func ConnectConsumer(kafkaBrokerUrl, topic string) *kafka.Reader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaBrokerUrl},
		GroupID:  "kafka_demo",
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	return r
}

func PushToQueue(kafkaBrokerUrl string, clientId string, topic string, message string) error {
	producer := ConnectProducer(kafkaBrokerUrl)
	defer producer.Close()

	err := producer.WriteMessages(context.Background(), kafka.Message{
		Topic: topic,
		Value: []byte(message),
	})
	if err != nil {
		return err
	}

	logrus.Infof("Message is stored in topic(%s)\n", topic)

	return nil
}
