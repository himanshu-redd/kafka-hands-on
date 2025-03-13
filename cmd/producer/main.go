package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	Longitude = 19.907397600347032
	Latitude  = 83.16382600762172
	ApiKey    = "ef46985d12600a2a12d7bb43dc7432ef"
)

func main() {
	fmt.Printf("Staring producer\n")

	resp, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&appid=%v",
		Longitude, Latitude, ApiKey))
	if err != nil {
		log.Printf("error: %v\n", err.Error())
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("Response: %v\n", string(body))
}
