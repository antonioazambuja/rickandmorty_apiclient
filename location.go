package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

const baseLocationURL string = "https://rickandmortyapi.com/api/location/"

// GetLocation - get location by ID
func GetLocation(id int) *Location {
	responseLocation, errLocation := http.Get(baseLocationURL + strconv.Itoa(id))
	if errLocation != nil {
		panic(errLocation)
	}
	LogLocation.Println("Get location by id")
	var location Location
	errDecode := json.NewDecoder(responseLocation.Body).Decode(&location)
	if errDecode != nil {
		panic(errDecode)
	}
	return &location
}

// GetAllLocations - get all locations
func GetAllLocations() *AllLocations {
	responseAllLocation, errAllLocation := http.Get(baseLocationURL)
	if errAllLocation != nil {
		panic(errAllLocation)
	}
	LogLocation.Println("Get all locations")
	var allLocation AllLocations
	errDecode := json.NewDecoder(responseAllLocation.Body).Decode(&allLocation)
	if errDecode != nil {
		panic(errDecode)
	}
	return &allLocation
}
