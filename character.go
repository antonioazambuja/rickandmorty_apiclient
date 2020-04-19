package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

const baseCharacterURL string = "https://rickandmortyapi.com/api/character/"

// GetCharacter - get character by ID
func GetCharacter(id int) Character {
	// LogCharacter.Printf("Get character by id: %d\n", id)
	responseCharacter, errCharacter := http.Get(baseCharacterURL + strconv.Itoa(id))
	if errCharacter != nil {
		panic(errCharacter)
	}
	var characterResponse CharacterResponse
	errDecode := json.NewDecoder(responseCharacter.Body).Decode(&characterResponse)
	if errDecode != nil {
		panic(errDecode)
	}
	return _ResponseToCharacter(characterResponse)
}

// GetMultipleCharacters - get multiple characters by array int
func GetMultipleCharacters(ids []int) []Character {
	// LogCharacter.Printf("Get multiple characters")
	var characters []Character
	for _, id := range ids {
		characters = append(characters, GetCharacter(id))
	}
	return characters
}

// GetAllCharacters - get all characters
func GetAllCharacters() []Character {
	var characters []Character
	// LogCharacter.Println("Get all characters")
	for i := 1; ; i++ {
		responseAllCharacter, errAllCharacter := http.Get(baseCharacterURL + "?page=" + strconv.Itoa(i))
		if errAllCharacter != nil {
			panic(errAllCharacter)
		}
		var allCharactersResponse AllCharacterResponse
		errDecode := json.NewDecoder(responseAllCharacter.Body).Decode(&allCharactersResponse)
		if errDecode != nil {
			panic(errDecode)
		}
		for _, character := range allCharactersResponse.Characters {
			// LogCharacter.Printf("Get character by id: %d\n", character.ID)
			characters = append(characters, _ResponseToCharacter(character))
		}
		if allCharactersResponse.Info.Next == "" {
			break
		}
	}
	return characters
}

// FilterCharacters - get filtered characters by parameter
func FilterCharacters(filterFunc func(character Character, comparation string) bool, comparation string) []Character {
	// LogCharacter.Println("Get filtered characters")
	var characters []Character
	allCharacters := GetAllCharacters()
	for _, character := range allCharacters {
		if filterFunc(character, comparation) {
			characters = append(characters, character)
		}
	}
	return characters
}

// _GetEpisodes - get characters of Character
func _GetEpisodesIDS(characterResponse CharacterResponse) []int {
	var episodes []int
	for index := range characterResponse.Episodes {
		episodes = append(episodes, index+1)
	}
	return episodes
}

func _ResponseToCharacter(characterResponse CharacterResponse) Character {
	return Character{
		ID:       characterResponse.ID,
		Name:     characterResponse.Name,
		Status:   characterResponse.Status,
		Species:  characterResponse.Species,
		Type:     characterResponse.Type,
		Gender:   characterResponse.Gender,
		Origin:   ParseURL(characterResponse.Origin.URL),
		Location: ParseURL(characterResponse.Location.URL),
		Image:    characterResponse.Image,
		Episodes: _GetEpisodesIDS(characterResponse),
		URL:      characterResponse.URL,
		Created:  characterResponse.Created,
	}
}
