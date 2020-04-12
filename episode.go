package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

const baseEpisodeURL string = "https://rickandmortyapi.com/api/episode/"

// GetEpisode - get episode by ID
func GetEpisode(id int) *Episode {
	responseEpisode, errEpisode := http.Get(baseEpisodeURL + strconv.Itoa(id))
	if errEpisode != nil {
		panic(errEpisode)
	}
	LogEpisode.Println("Get episode by id")
	var episode Episode
	errDecode := json.NewDecoder(responseEpisode.Body).Decode(&episode)
	if errDecode != nil {
		panic(errDecode)
	}
	return &episode
}

// GetAllEpisodes - get all episodes
func GetAllEpisodes() *AllEpisodes {
	responseAllEpisodes, errAllEpisodes := http.Get(baseEpisodeURL)
	if errAllEpisodes != nil {
		panic(errAllEpisodes)
	}
	LogEpisode.Println("Get all episodes")
	var allEpisodes AllEpisodes
	errDecode := json.NewDecoder(responseAllEpisodes.Body).Decode(&allEpisodes)
	if errDecode != nil {
		panic(errDecode)
	}
	return &allEpisodes
}
