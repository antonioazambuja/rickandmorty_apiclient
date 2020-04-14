package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

const baseLocationURL string = "https://rickandmortyapi.com/api/location/"

// GetLocation - get location by ID
func GetLocation(id int) Location {
	responseLocation, errLocation := http.Get(baseLocationURL + strconv.Itoa(id))
	if errLocation != nil {
		panic(errLocation)
	}
	LogLocation.Printf("Get location by id: %d\n", id)
	var locationResponse LocationResponse
	errDecode := json.NewDecoder(responseLocation.Body).Decode(&locationResponse)
	if errDecode != nil {
		panic(errDecode)
	}
	return _ResponseToLocation(locationResponse)
}

// GetAllLocations - get all locations
func GetAllLocations() []Location {
	LogLocation.Println("Get all locations")
	var locations []Location
	for i := 1; ; i++ {
		responseAllLocation, errAllLocation := http.Get(baseLocationURL + "?page=" + strconv.Itoa(i))
		if errAllLocation != nil {
			panic(errAllLocation)
		}
		var allLocations AllLocationResponse
		errDecode := json.NewDecoder(responseAllLocation.Body).Decode(&allLocations)
		if errDecode != nil {
			panic(errDecode)
		}
		for _, location := range allLocations.Locations {
			locations = append(locations, _ResponseToLocation(location))
		}
		if allLocations.Info.Next == "" {
			break
		}
	}
	return locations
}

// _GetResidents - get characters residents of Location
func _GetResidents(locationResponse LocationResponse) []int {
	var characters []int
	for index := range locationResponse.Residents {
		characters = append(characters, index+1)
	}
	return characters
}

func _ResponseToLocation(locationResponse LocationResponse) Location {
	return Location{
		ID:        locationResponse.ID,
		Name:      locationResponse.Name,
		Type:      locationResponse.Type,
		Dimension: locationResponse.Dimension,
		Residents: _GetResidents(locationResponse),
		Created:   locationResponse.Created,
		URL:       locationResponse.URL,
	}
}
