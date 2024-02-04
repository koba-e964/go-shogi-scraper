package scraper

type Player struct {
	Kind string `json:"kind"` // "pro" or "lady"
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

type PlayerInfo struct {
	ID            int    `json:"id"`
	URL           string `json:"url"`
	RetrievalTime string `json:"retrieval_time"`
	HashAlgorithm string `json:"hash_algorithm"`
	Hash          []byte `json:"hash"`
	NameJP        string `json:"name_jp"`
	NameEN        string `json:"name_en"`
	Birthday      string `json:"birthday"`
	Birthplace    string `json:"birthplace"`
	Mentor        Player `json:"mentor"`
}

type RetrievedData interface {
	GetURL() string
	GetRetrievalTime() string
}
