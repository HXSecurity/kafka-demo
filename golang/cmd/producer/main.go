package main

import (
	"net/http"
	"os"

	_ "github.com/HXSecurity/DongTai-agent-go/run/base"
	_ "github.com/HXSecurity/DongTai-agent-go/run/http"
	"github.com/sirupsen/logrus"

	"kafkademo/kafka"
)

func main() {
	// connect to kafka: x.x.x.x:9092
	kafkaBroker := []string{os.Getenv("KAFKA_BROKER")}
	kafkaProducer, errConnection := kafka.ConnectProducer(kafkaBroker)
	if errConnection != nil {
		logrus.Printf("error: %s", "Unable to configure kafka")
		return
	}
	defer kafkaProducer.Close()

	kafkaClientId := os.Getenv("KAFKA_CLIENT")
	kafkaTopic := os.Getenv("KAFKA_TOPIC")
	if kafkaTopic == "" {
		kafkaTopic = "addUserV2"
	}

	http.HandleFunc("/kafka/publish", func(writer http.ResponseWriter, request *http.Request) {
		values := request.URL.Query()
		message := values.Get("message")
		if message == "" {
			response(writer, "parameter message is required")
			return
		}

		err := kafka.PushToQueue(kafkaBroker, kafkaClientId, kafkaTopic, message)
		if err != nil {
			logrus.Errorf("error while push message into kafka: %s", err)
			response(writer, "failed")
			return
		}

		response(writer, "ok")
	})

	logrus.Infoln("listening on port 8811")
	err := http.ListenAndServe(":8811", nil)
	if err != nil {
		logrus.Errorf("http listen failed: %s", err)
	}
}

func response(writer http.ResponseWriter, content string) {
	_, err := writer.Write([]byte(content))
	if err != nil {
		logrus.Errorf("response failed: %s", err)
	}
}
