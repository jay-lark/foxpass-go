package main

import (
	"io/ioutil"
	"log"
	"net/http"
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
	if err != nil {
		log.Fatalf("error reading HTTP response body: %v", err)
	}

	log.Println("We got the response:", string(responseBytes))
}
