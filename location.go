package rickandmortyapiclient

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

const baseLocationURL string = "https://rickandmortyapi.com/api/location/"

// GetLocation - get location by ID
func GetLocation(id int) (Location, error) {
	responseLocation, errLocation := http.Get(baseLocationURL + strconv.Itoa(id))
	if errLocation != nil || responseLocation.StatusCode != 200 {
		return Location{}, errors.New("Location API: error get location")
	}
	var locationResponse LocationResponse
	errDecode := json.NewDecoder(responseLocation.Body).Decode(&locationResponse)
	if errDecode != nil {
		return Location{}, errDecode
	}
	return _ResponseToLocation(locationResponse), nil
}

// GetAllLocations - get all locations
func GetAllLocations() ([]Location, error) {
	var locations []Location
	for i := 1; ; i++ {
		responseAllLocation, errAllLocation := http.Get(baseLocationURL + "?page=" + strconv.Itoa(i))
		if errAllLocation != nil {
			return []Location{}, errAllLocation
		}
		var allLocations AllLocationResponse
		errDecode := json.NewDecoder(responseAllLocation.Body).Decode(&allLocations)
		if errDecode != nil {
			return []Location{}, errDecode
		}
		for _, location := range allLocations.Locations {
			locations = append(locations, _ResponseToLocation(location))
		}
		if allLocations.Info.Next == "" {
			break
		}
	}
	return locations, nil
}

// FilterLocations - get filtered locations by parameter
func FilterLocations(filterFunc func(location Location, comparation string) bool, comparation string) ([]Location, error) {
	var locations []Location
	allLocations, errAllLocation := GetAllLocations()
	if errAllLocation != nil {
		return []Location{}, errAllLocation
	}
	for _, location := range allLocations {
		if filterFunc(location, comparation) {
			locations = append(locations, location)
		}
	}
	return locations, nil
}

// GetMultipleLocations - get multiple locations by array int
func GetMultipleLocations(ids ...int) ([]Location, error) {
	var locations []Location
	for _, id := range ids {
		location, err := GetLocation(id)
		if err != nil {
			return []Location{}, err
		}
		locations = append(locations, location)
	}
	return locations, nil
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
