package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/IBM/sarama"
	"gocloud.dev/pubsub"
	"gocloud.dev/pubsub/kafkapubsub"
)

const (
	Longitude = 19.907397600347032
	Latitude  = 83.16382600762172
	ApiKey    = "ef46985d12600a2a12d7bb43dc7432ef"
)

func main() {
	log.Printf("Staring producer\n")

	resp, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&appid=%v",
		Longitude, Latitude, ApiKey))
	if err != nil {
		log.Printf("error: %v\n", err.Error())
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Weather Response: %v\n", string(body))

	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.Timeout = 5
	saramaConfig.Producer.Transaction.Retry.Max = 20
	saramaConfig.Producer.Return.Successes = true

	topic, err := kafkapubsub.OpenTopic([]string{"localhost:9092"}, saramaConfig, "current-weather", nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	bytes := []byte(string(body))
	err = topic.Send(context.Background(), &pubsub.Message{
		Body: bytes,
	})

	if err != nil {
		log.Fatal(err.Error())
	} else {
		log.Printf("Message sent successfully\n")
	}
}
