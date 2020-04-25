package rickandmortyapiclient

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEpisode(test *testing.T) {
	_, errorTest1 := GetEpisode(1)
	assert.Empty(test, errorTest1)
	_, errorTest2 := GetEpisode(01)
	assert.Empty(test, errorTest2)
	_, errorTest3 := GetEpisode(1.0)
	assert.Empty(test, errorTest3)
	_, errorTest4 := GetEpisode(0)
	assert.NotEmpty(test, errorTest4)
	_, errorTest5 := GetEpisode(-1)
	assert.NotEmpty(test, errorTest5)
}

func TestGetMultipleEpisodes(test *testing.T) {
	_, errorTest1 := GetMultipleEpisodes(1, 2, 3)
	assert.Empty(test, errorTest1)
	_, errorTest2 := GetMultipleEpisodes(0, 1, 2, 3)
	assert.NotEmpty(test, errorTest2)
	_, errorTest3 := GetMultipleEpisodes(10, 01, 02, 030)
	assert.Empty(test, errorTest3)
	_, errorTest4 := GetMultipleEpisodes(-1, 10, 01, 02, 030)
	assert.NotEmpty(test, errorTest4)
}

func TestGetAllEpisodes(test *testing.T) {
	_, errorTest1 := GetAllEpisodes()
	assert.Empty(test, errorTest1)
}

func TestFilterEpisodes(test *testing.T) {
	byID := func(episode Episode, comparation string) bool {
		id, err := strconv.Atoi(comparation)
		if err != nil {
			fmt.Print(err)
		}
		if episode.ID == id {
			return true
		}
		return false
	}
	_, errorFilterEpisodesTest1 := FilterEpisodes(byID, "10")
	assert.Empty(test, errorFilterEpisodesTest1)

	byName := func(episode Episode, name string) bool {
		return strings.Contains(episode.Name, name)
	}
	_, errorFilterEpisodesTest2 := FilterEpisodes(byName, "Rick")
	assert.Empty(test, errorFilterEpisodesTest2)

	byEpisodeName := func(episode Episode, name string) bool {
		return strings.Contains(episode.Episode, name)
	}
	_, errorFilterEpisodesTest3 := FilterEpisodes(byEpisodeName, "S03")
	assert.Empty(test, errorFilterEpisodesTest3)
}
