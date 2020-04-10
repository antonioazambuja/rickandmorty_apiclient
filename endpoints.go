package main

import (
	"encoding/json"
	"net/http"
)

const baseEndpoint string = "https://rickandmortyapi.com/api/"

func getEndpoints() map[string]string {
	response, errGet := http.Get(baseEndpoint)
	if errGet != nil {
		panic(errGet)
	}
	defer response.Body.Close()
	// body, errBody := ioutil.ReadAll(response.Body)
	// if errBody != nil {
	// 	panic(errBody)
	// }
	// var endpoints []Endpoint
	var result map[string]string
	errDecode := json.NewDecoder(response.Body).Decode(&result)

	// fmt.Print(string(body))
	// errDecode := json.Unmarshal(response.Body), &result)
	if errDecode != nil {
		panic(errDecode)
	}
	return result
}
