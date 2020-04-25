# Rick and Morty API Client

## Description

This project is based in API of create for [Axel Fuhrmann](https://axelfuhrmann.com/).

Check for more informations about API: [Rick And Morty](https://rickandmortyapi.com/).

## Usage

```
package main

import (
	"fmt"
	"strconv"
)

func main() {
	character, errGetCharater := GetCharacter(1)
	if errGetCharater != nil {
		panic(errGetCharater)
	}
	fmt.Println(character)

	allCharacters, errGetLocation := GetAllCharacters()
	if errGetLocation != nil {
		panic(errGetLocation)
	}
	fmt.Println(allCharacters)

	multiplesCharacters, errGetLocation := GetMultipleCharacters(1,10,29,101)
	if errGetLocation != nil {
		panic(errGetLocation)
	}
	fmt.Println(multiplesCharacters)

	byID := func (character Character, comparation string) bool {
		id, err := strconv.Atoi(comparation)
		if err != nil {
			panic(err)
		}
		if character.ID == id {
			return true
		}
		return false
	}
	filterCharacters, errGetLocation := FilterCharacters(byID, "10")
	if errGetLocation != nil {
		panic(errGetLocation)
	}
	fmt.Println(filterCharacters)
}

```

These functions are available for all schemes present in the API.

### Thanks!