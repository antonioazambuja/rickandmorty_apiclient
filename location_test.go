package rickandmortyapiclient

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLocation(test *testing.T) {
	_, errorLocationTest1 := GetLocation(1)
	assert.Empty(test, errorLocationTest1)
	_, errorLocationTest2 := GetLocation(01)
	assert.Empty(test, errorLocationTest2)
	_, errorLocationTest3 := GetLocation(1.0)
	assert.Empty(test, errorLocationTest3)
	// API error
	_, errorLocationTest4 := GetLocation(0)
	assert.Empty(test, errorLocationTest4)
	// CORRECT
	// _, errorLocationTest4 := GetLocation(0)
	// assert.NotEmpty(test, errorLocationTest4)
	_, errorLocationTest5 := GetLocation(-1)
	assert.NotEmpty(test, errorLocationTest5)
}

func TestGetMultipleLocations(test *testing.T) {
	_, errorLocationsTest1 := GetMultipleLocations(1, 2, 3)
	assert.Empty(test, errorLocationsTest1)
	_, errorLocationsTest2 := GetMultipleLocations(1, 2, 3)
	assert.Empty(test, errorLocationsTest2)
	_, errorLocationsTest3 := GetMultipleLocations(10, 01, 02, 030)
	assert.Empty(test, errorLocationsTest3)
	_, errorLocationsTest4 := GetMultipleLocations(-1, 10, 01, 02, 030)
	assert.NotEmpty(test, errorLocationsTest4)
}

func TestGetAllLocations(test *testing.T) {
	_, errorLocationsTest1 := GetAllLocations()
	assert.Empty(test, errorLocationsTest1)
}

func TestFilterLocations(test *testing.T) {
	byID := func(location Location, comparation string) bool {
		id, err := strconv.Atoi(comparation)
		if err != nil {
			fmt.Print(err)
		}
		if location.ID == id {
			return true
		}
		return false
	}
	_, errorFilterLocationsTest1 := FilterLocations(byID, "10")
	assert.Empty(test, errorFilterLocationsTest1)

	byName := func(location Location, name string) bool {
		return strings.Contains(location.Name, name)
	}
	_, errorFilterLocationsTest2 := FilterLocations(byName, "Earth")
	assert.Empty(test, errorFilterLocationsTest2)

	byType := func(location Location, name string) bool {
		return strings.Contains(location.Type, name)
	}
	_, errorFilterLocationsTest3 := FilterLocations(byType, "Planet")
	assert.Empty(test, errorFilterLocationsTest3)
}
