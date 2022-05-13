package main

import (
	"net/http"
	"os"

	_ "github.com/HXSecurity/DongTai-agent-go/run/base"
	_ "github.com/HXSecurity/DongTai-agent-go/run/http"
	_ "github.com/HXSecurity/DongTai-agent-go/run/kafkaGo"
	"github.com/sirupsen/logrus"

	"kafkademo/kafka_go"
)

func main() {
	// connect to kafka: x.x.x.x:9092
	kafkaBroker := os.Getenv("KAFKA_BROKER")
	kafkaProducer := kafka_go.ConnectProducer(kafkaBroker)
	defer kafkaProducer.Close()

	kafkaClientId := os.Getenv("KAFKA_CLIENT")

	http.HandleFunc("/kafka/publish", func(writer http.ResponseWriter, request *http.Request) {
		values := request.URL.Query()

		kafkaTopic := values.Get("topic")
		if kafkaTopic == "" {
			kafkaTopic = "addUserV2"
		}

		message := values.Get("message")
		if message == "" {
			response(writer, "parameter message is required")
			return
		}

		err := kafka_go.PushToQueue(kafkaBroker, kafkaClientId, kafkaTopic, message)
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
