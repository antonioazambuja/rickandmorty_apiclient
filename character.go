package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const baseCharacterURL string = "https://rickandmortyapi.com/api/character/"

// GetAllCharacters - get all characters of API
func GetAllCharacters() {
	LogCharacter.Printf("Get all characters")
	responseGetAllCharacter, errGetAllCharacter := http.Get(baseCharacterURL)
	if errGetAllCharacter != nil {
		panic(errGetAllCharacter)
	}
	defer responseGetAllCharacter.Body.Close()
	var result AllCharacters
	errDecode := json.NewDecoder(responseGetAllCharacter.Body).Decode(&result)
	if errDecode != nil {
		panic(errDecode)
	}
	fmt.Println(result.Info)
}

// GetCharacter - get character by id
func GetCharacter(id int) *Character {
	LogCharacter.Printf("Get character of ID: %d", id)
	responseGetCharacter, errGetCharacter := http.Get(baseCharacterURL + strconv.Itoa(id))
	if errGetCharacter != nil {
		panic(errGetCharacter)
	}
	defer responseGetCharacter.Body.Close()
	var character Character
	errDecode := json.NewDecoder(responseGetCharacter.Body).Decode(&character)
	if errDecode != nil {
		panic(errDecode)
	}
	return &character
}
