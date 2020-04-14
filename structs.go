package main

// Info - info get any all schema
type Info struct {
	Count int    `json:"count,omitempty"`
	Pages int    `json:"pages,omitempty"`
	Next  string `json:"next,omitempty"`
	Prev  string `json:"prev,omitempty"`
}

// Origin - origin of character
type Origin struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

// LocationCharacter - location of character
type LocationCharacter struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

// AllLocationResponse is struct of all locations
type AllLocationResponse struct {
	Info      Info               `json:"info"`
	Locations []LocationResponse `json:"results"`
}

// LocationResponse - is struct of location in API
type LocationResponse struct {
	ID        int      `json:"id,omitempty"`
	Name      string   `json:"name,omitempty"`
	Type      string   `json:"type,omitempty"`
	Dimension string   `json:"dimension,omitempty"`
	Residents []string `json:"residents,omitempty"`
	URL       string   `json:"url,omitempty"`
	Created   string   `json:"string,omitempty"`
}

// Location - locations availables
type Location struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Type      string `json:"type,omitempty"`
	Dimension string `json:"dimension,omitempty"`
	Residents []int  `json:"residents,omitempty"`
	URL       string `json:"url,omitempty"`
	Created   string `json:"string,omitempty"`
}

// CharacterResponse is struct of character in API
type CharacterResponse struct {
	ID       int               `json:"id,omitempty"`
	Name     string            `json:"name,omitempty"`
	Status   string            `json:"status,omitempty"`
	Species  string            `json:"species,omitempty"`
	Type     string            `json:"type,omitempty"`
	Gender   string            `json:"gender,omitempty"`
	Origin   Origin            `json:"origin,omitempty"`
	Location LocationCharacter `json:"location,omitempty"`
	Image    string            `json:"image,omitempty"`
	Episodes []string          `json:"episode,omitempty"`
	URL      string            `json:"url,omitempty"`
	Created  string            `json:"created,omitempty"`
}

// AllCharacterResponse is struct of all characters
type AllCharacterResponse struct {
	Info       Info                `json:"info"`
	Characters []CharacterResponse `json:"results"`
}

// Character is struct of character in API
type Character struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Status   string `json:"status,omitempty"`
	Species  string `json:"species,omitempty"`
	Type     string `json:"type,omitempty"`
	Gender   string `json:"gender,omitempty"`
	Origin   int    `json:"origin,omitempty"`
	Location int    `json:"location,omitempty"`
	Image    string `json:"image,omitempty"`
	Episodes []int  `json:"episode,omitempty"`
	URL      string `json:"url,omitempty"`
	Created  string `json:"created,omitempty"`
}

// EpisodeResponse is struct of episode in API
type EpisodeResponse struct {
	ID         int      `json:"id,omitempty"`
	Name       string   `json:"name,omitempty"`
	Airdate    string   `json:"air_date,omitempty"`
	Episode    string   `json:"episode,omitempty"`
	Characters []string `json:"characters,omitempty"`
	URL        string   `json:"url,omitempty"`
	Created    string   `json:"created,omitempty"`
}

// AllEpisodeResponse is struct of all characters
type AllEpisodeResponse struct {
	Info     Info              `json:"info"`
	Episodes []EpisodeResponse `json:"results"`
}

// Episode is struct of episode
type Episode struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Airdate    string `json:"air_date,omitempty"`
	Episode    string `json:"episode,omitempty"`
	Characters []int  `json:"characters,omitempty"`
	URL        string `json:"url,omitempty"`
	Created    string `json:"created,omitempty"`
}
