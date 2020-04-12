package main

// Origin - origin of character
type Origin struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

// Location - last location of character
type Location struct {
	ID        int      `json:"id,omitempty"`
	Name      string   `json:"name,omitempty"`
	Types     string   `json:"type,omitempty"`
	Dimension string   `json:"dimension,omitempty"`
	Residents []string `json:"residents,omitempty"`
	URL       string   `json:"url,omitempty"`
	Created   string   `json:"string,omitempty"`
}

// Info - info get any all schema
type Info struct {
	Count int    `json:"count,omitempty"`
	Pages int    `json:"pages,omitempty"`
	Next  string `json:"next,omitempty"`
	Prev  string `json:"prev,omitempty"`
}

// AllLocations is struct of all locations
type AllLocations struct {
	Info    Info       `json:"info"`
	Results []Location `json:"results"`
}

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

// AllCharacters is struct of all characters
type AllCharacters struct {
	Info    Info        `json:"info"`
	Results []Character `json:"results"`
}
