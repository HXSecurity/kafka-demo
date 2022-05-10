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
