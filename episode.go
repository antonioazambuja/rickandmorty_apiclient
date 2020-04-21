package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

const baseEpisodeURL string = "https://rickandmortyapi.com/api/episode/"

// GetEpisode - get episode by ID
func GetEpisode(id int) Episode {
	responseEpisode, errEpisode := http.Get(baseEpisodeURL + strconv.Itoa(id))
	if errEpisode != nil {
		panic(errEpisode)
	}
	var episodeResponse EpisodeResponse
	errDecode := json.NewDecoder(responseEpisode.Body).Decode(&episodeResponse)
	if errDecode != nil {
		panic(errDecode)
	}
	return _ResponseToEpisode(episodeResponse)
}

// GetAllEpisodes - get all episodes
func GetAllEpisodes() []Episode {
	var episodes []Episode
	for i := 1; ; i++ {
		responseAllEpisodes, errAllEpisodes := http.Get(baseEpisodeURL + "?page=" + strconv.Itoa(i))
		if errAllEpisodes != nil {
			panic(errAllEpisodes)
		}
		var allEpisodes AllEpisodeResponse
		errDecode := json.NewDecoder(responseAllEpisodes.Body).Decode(&allEpisodes)
		if errDecode != nil {
			panic(errDecode)
		}
		for _, episode := range allEpisodes.Episodes {
			episodes = append(episodes, _ResponseToEpisode(episode))
		}
		if allEpisodes.Info.Next == "" {
			break
		}
	}
	return episodes
}

// FilterEpisodes - get filtered episodes by parameter
func FilterEpisodes(filterFunc func(episode Episode, comparation string) bool, comparation string) []Episode {
	var episodes []Episode
	allEpisodes := GetAllEpisodes()
	for _, episode := range allEpisodes {
		if filterFunc(episode, comparation) {
			episodes = append(episodes, episode)
		}
	}
	return episodes
}

// GetMultipleEpisodes - get multiple episodes by array int
func GetMultipleEpisodes(ids []int) []Episode {
	var episodes []Episode
	for _, id := range ids {
		episodes = append(episodes, GetEpisode(id))
	}
	return episodes	
}

// _GetCharacters - get characters of Episode
func _GetCharacters(episodeResponse EpisodeResponse) []int {
	var characters []int
	for index := range episodeResponse.Characters {
		characters = append(characters, index+1)
	}
	return characters
}

func _ResponseToEpisode(episodeResponse EpisodeResponse) Episode {
	var episode Episode
	episode.Airdate = episodeResponse.Airdate
	episode.Characters = _GetCharacters(episodeResponse)
	episode.Created = episodeResponse.Created
	episode.Episode = episodeResponse.Episode
	episode.ID = episodeResponse.ID
	episode.Name = episodeResponse.Name
	episode.URL = episodeResponse.URL
	return episode
}
