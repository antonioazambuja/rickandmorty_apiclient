package main

// Character is struct of character in API
type Character struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Status     string `json:"status"`
	Species    string `json:"species"`
	Subspecies string `json:"type"`
	Gender     string `json:"gender"`
	Origin     string `json:"origin"`
	Location   string `json:"location"`
	Image      string `json:"image"`
	Episode    string `json:"episode"`
	URL        string `json:"url"`
	Created    string `json:"created"`
}
