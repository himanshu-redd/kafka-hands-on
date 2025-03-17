package main

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"gocloud.dev/pubsub/kafkapubsub"
)

func main() {

	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.Timeout = 5
	saramaConfig.Producer.Transaction.Retry.Max = 20
	saramaConfig.Producer.Return.Successes = true

	topic, err := kafkapubsub.OpenSubscription([]string{"localhost:9092"},
		saramaConfig,
		"fs",
		[]string{"current-weather"},
		nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	msg, err := topic.Receive(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Message received: ", string(msg.Body))

}
