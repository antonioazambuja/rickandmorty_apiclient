package rickandmortyapiclient

import (
	"fmt"
	"strconv"
)

func byID(character Character, comparation string) bool {
	id, err := strconv.Atoi(comparation)
	if err != nil {
		panic(err)
	}
	if character.ID == id {
		return true
	}
	return false
}

func main() {
	fmt.Println(FilterCharacters(byID, "10"))
}
