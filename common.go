package scraper

type Player struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

type PlayerInfo struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	URL        string `json:"url"`
	Birthday   string `json:"birthday"`
	Birthplace string `json:"birthplace"`
	Mentor     Player `json:"mentor"`
}
