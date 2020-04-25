package rickandmortyapiclient

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

const baseEpisodeURL string = "https://rickandmortyapi.com/api/episode/"

// GetEpisode - get episode by ID
func GetEpisode(id int) (Episode, error) {
	responseEpisode, errEpisode := http.Get(baseEpisodeURL + strconv.Itoa(id))
	if errEpisode != nil || responseEpisode.StatusCode != 200 {
		return Episode{}, errors.New("Episode API: error get episode")
	}
	var episodeResponse EpisodeResponse
	errDecode := json.NewDecoder(responseEpisode.Body).Decode(&episodeResponse)
	if errDecode != nil {
		return Episode{}, errDecode
	}
	return _ResponseToEpisode(episodeResponse), nil
}

// GetAllEpisodes - get all episodes
func GetAllEpisodes() ([]Episode, error) {
	var episodes []Episode
	for i := 1; ; i++ {
		responseAllEpisodes, errAllEpisodes := http.Get(baseEpisodeURL + "?page=" + strconv.Itoa(i))
		if errAllEpisodes != nil {
			return []Episode{}, errAllEpisodes
		}
		var allEpisodes AllEpisodeResponse
		errDecode := json.NewDecoder(responseAllEpisodes.Body).Decode(&allEpisodes)
		if errDecode != nil {
			return []Episode{}, errDecode
		}
		for _, episode := range allEpisodes.Episodes {
			episodes = append(episodes, _ResponseToEpisode(episode))
		}
		if allEpisodes.Info.Next == "" {
			break
		}
	}
	return episodes, nil
}

// FilterEpisodes - get filtered episodes by parameter
func FilterEpisodes(filterFunc func(episode Episode, comparation string) bool, comparation string) ([]Episode, error) {
	var episodes []Episode
	allEpisodes, errorGetAllEpisodes := GetAllEpisodes()
	if errorGetAllEpisodes != nil {
		return []Episode{}, errorGetAllEpisodes
	}
	for _, episode := range allEpisodes {
		if filterFunc(episode, comparation) {
			episodes = append(episodes, episode)
		}
	}
	return episodes, nil
}

// GetMultipleEpisodes - get multiple episodes by array int
func GetMultipleEpisodes(ids ...int) ([]Episode, error) {
	var episodes []Episode
	for _, id := range ids {
		episode, err := GetEpisode(id)
		if err != nil {
			return []Episode{}, err
		}
		episodes = append(episodes, episode)
	}
	return episodes, nil
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
