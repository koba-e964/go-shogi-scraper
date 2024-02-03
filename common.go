package scraper

type Player struct {
	Kind string `json:"kind"` // "pro" or "lady"
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

type PlayerInfo struct {
	ID         int    `json:"id"`
	NameJP     string `json:"name_jp"`
	NameEN     string `json:"name_en"`
	URL        string `json:"url"`
	Birthday   string `json:"birthday"`
	Birthplace string `json:"birthplace"`
	Mentor     Player `json:"mentor"`
}
