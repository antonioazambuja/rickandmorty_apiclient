package main

import (
	"fmt"
)

func main() {
	// episode, err := json.Marshal(GetEpisode(5))
	// if err != nil {
	// 	panic(err)
	// }
	// characterIndent, err := json.MarshalIndent(GetCharacter(1), "", "  ")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(characterIndent))
	// characters := GetAllCharacters()
	location := GetAllLocations()
	fmt.Println(len(location))
}
