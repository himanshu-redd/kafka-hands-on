package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
	"gocloud.dev/pubsub/kafkapubsub"

	"github.com/himanshu-redd/kafka-hands-on/entities"
)

func main() {

	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.Timeout = 5
	saramaConfig.Producer.Transaction.Retry.Max = 20
	saramaConfig.Producer.Return.Successes = false

	topic, err := kafkapubsub.OpenSubscription([]string{"localhost:9092"},
		saramaConfig,
		"fs",
		[]string{"current-weather"},
		nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	for {
		msg, err := topic.Receive(context.Background())

		if err != nil {
			log.Fatal(err.Error())
		}

		if msg != nil {
			var weather entities.WeatherResponse
			err := json.Unmarshal(msg.Body, &weather)
			if err != nil {
				log.Fatalf("json ummarshal failed : %v\n", err.Error())
			}
			log.Printf("Current weather: %v\n", weather.Weather[0].Description)
		}
	}

}
