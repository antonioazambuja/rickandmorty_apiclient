package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Character is struct of character in API
type Character struct {
	ID         int      `json:"id,omitempty"`
	Name       string   `json:"name,omitempty"`
	Status     string   `json:"status,omitempty"`
	Species    string   `json:"species,omitempty"`
	Subspecies string   `json:"type,omitempty"`
	Gender     string   `json:"gender,omitempty"`
	Origin     Origin   `json:"origin,omitempty"`
	Location   Location `json:"location,omitempty"`
	Image      string   `json:"image,omitempty"`
	Episode    []string `json:"episode,omitempty"`
	URL        string   `json:"url,omitempty"`
	Created    string   `json:"created,omitempty"`
}

// Info - info get all character
type Info struct {
	Count int    `json:"count,omitempty"`
	Pages int    `json:"pages,omitempty"`
	Next  string `json:"next,omitempty"`
	Prev  string `json:"prev,omitempty"`
}

// AllCharacters is struct of all characters
type AllCharacters struct {
	Info    Info        `json:"info"`
	Results []Character `json:"results"`
}

func init() {
	response, err := http.Get(BaseEndpoint + "character/")
	if err != nil {
		panic(err)
	} else if response.StatusCode == 200 {
		LogCharacter.Println("Character endpoint is working!")
	}
}

// GetAllCharacters - get all characters of API
func GetAllCharacters() {
	LogCharacter.Printf("Get all characters")
	responseAllGetCharacter, errAllGetCharacter := http.Get(BaseEndpoint + "character/")
	if errAllGetCharacter != nil {
		panic(errAllGetCharacter)
	}
	defer responseAllGetCharacter.Body.Close()
	var result AllCharacters
	errDecode := json.NewDecoder(responseAllGetCharacter.Body).Decode(&result)
	if errDecode != nil {
		panic(errDecode)
	}
	fmt.Println(result.Info)
}

// GetCharacter - get character by id
func GetCharacter(id int) *Character {
	LogCharacter.Printf("Get character of ID: %d", id)
	responseGetCharacter, errGetCharacter := http.Get(BaseEndpoint + "character/" + strconv.Itoa(id))
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
