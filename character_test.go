package rickandmortyapiclient

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCharacter(test *testing.T) {
	_, errorCharacter1 := GetCharacter(1)
	assert.Empty(test, errorCharacter1)
	_, errorCharacter2 := GetCharacter(01)
	assert.Empty(test, errorCharacter2)
	_, errorCharacter3 := GetCharacter(1.0)
	assert.Empty(test, errorCharacter3)
	_, errorCharacter4 := GetCharacter(0)
	assert.NotEmpty(test, errorCharacter4)
	_, errorCharacter5 := GetCharacter(-1)
	assert.NotEmpty(test, errorCharacter5)
}

func TestGetMultipleCharacters(test *testing.T) {
	_, errorCharacters1 := GetMultipleCharacters(1, 2, 3)
	assert.Empty(test, errorCharacters1)
	_, errorCharacters2 := GetMultipleCharacters(0, 1, 2, 3)
	assert.NotEmpty(test, errorCharacters2)
	_, errorCharacters3 := GetMultipleCharacters(10, 01, 02, 030)
	assert.Empty(test, errorCharacters3)
	_, errorCharacters4 := GetMultipleCharacters(-1, 10, 01, 02, 030)
	assert.NotEmpty(test, errorCharacters4)
}

func TestGetAllCharacters(test *testing.T) {
	_, errorAllCharacters := GetAllCharacters()
	assert.Empty(test, errorAllCharacters)
}

func TestFilterCharacters(test *testing.T) {
	byID := func(character Character, comparation string) bool {
		id, err := strconv.Atoi(comparation)
		if err != nil {
			fmt.Print(err)
		}
		if character.ID == id {
			return true
		}
		return false
	}
	_, errorCharactersTest1 := FilterCharacters(byID, "10")
	assert.Empty(test, errorCharactersTest1)

	byName := func(character Character, name string) bool {
		return strings.Contains(character.Name, name)
	}
	_, errorCharactersTest2 := FilterCharacters(byName, "Rick")
	assert.Empty(test, errorCharactersTest2)

	byGender := func(character Character, gender string) bool {
		return strings.Contains(character.Gender, gender)
	}
	_, errorCharactersTest3 := FilterCharacters(byGender, "Male")
	assert.Empty(test, errorCharactersTest3)
}
