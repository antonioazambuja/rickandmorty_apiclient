package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Origin - origin of character
type Origin struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

// Location - last location of character
type Location struct {
	ID        int      `json:"id,omitempty"`
	Name      string   `json:"name,omitempty"`
	Types     string   `json:"type,omitempty"`
	Dimension string   `json:"dimension,omitempty"`
	Residents []string `json:"residents,omitempty"`
	URL       string   `json:"url,omitempty"`
	Created   string   `json:"string,omitempty"`
}

// CreateOriginCharacter - create object origin of character
func CreateOriginCharacter(name, url string) *Origin {
	return &Origin{Name: name, URL: url}
}

// GetLocation - get location by ID
func GetLocation(id int) *Location {
	responseLocation, errLocation := http.Get(BaseEndpoint + "/location/" + strconv.Itoa(id))
	if errLocation != nil {
		panic(errLocation)
	}
	var location Location
	errDecode := json.NewDecoder(responseLocation.Body).Decode(&location)
	if errDecode != nil {
		panic(errDecode)
	}
	return &location
}

// GetLocationByCharacter - create object location of character
func GetLocationByCharacter(name, url string) *Location {
	responseLocation, errLocation := http.Get(url)
	if errLocation != nil {
		panic(errLocation)
	}
	defer responseLocation.Body.Close()
	var location Location
	errDecode := json.NewDecoder(responseLocation.Body).Decode(&location)
	if errDecode != nil {
		panic(errDecode)
	}
	return &location
}
