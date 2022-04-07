package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jay-lark/foxpass-go/foxpass"
)

func main() {
	req, err := http.NewRequest(
		http.MethodGet,
		"https://api.foxpass.com/v1/users",
		nil,
	)
	if err != nil {
		log.Fatalf("error creating HTTP request: %v", err)
	}

	req.Header.Add("Authorization", "Token "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("error sending HTTP request: %v", err)
	}
	responseBytes, err := ioutil.ReadAll(res.Body)
	var weatherSamples []foxpass.Users
	if err := json.Unmarshal(responseBytes, &weatherSamples); err != nil {
		log.Fatalf("error deserializing weather data")
	}

	for _, w := range foxpass.Users {
		if w.Temp != nil && w.Temp.Value != nil {
			log.Printf("The temperature at %s \n",
				w.Email.Value)
		} else {
			log.Printf("No temperature data available at \n")
		}
	}
	if err != nil {
		log.Fatalf("error reading HTTP response body: %v", err)
	}

	log.Println("We got the response:", string(responseBytes))
}
