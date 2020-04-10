package main

import (
	"encoding/json"
	"net/http"
)

// BaseEndpoint - endpoint of API
const BaseEndpoint string = "https://rickandmortyapi.com/api/"

func init() {
	response, errGet := http.Get(BaseEndpoint)
	if errGet != nil {
		panic(errGet)
	}
	defer response.Body.Close()
	var result map[string]string
	errDecode := json.NewDecoder(response.Body).Decode(&result)
	if errDecode != nil {
		panic(errDecode)
	}
	for typeURL, url := range result {
		response, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		if response.StatusCode != 200 {
			panic("Endpoint " + typeURL + " unavailable!")
		}
		LogDefault.Printf("Endpoint %s OK!\n", typeURL)
	}
}
