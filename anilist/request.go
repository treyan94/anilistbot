package anilist

type Request struct {
	Query           string `json:"query"`
	SearchVariables `json:"variables"`
}

type SearchVariables struct {
	Search  string `json:"search"`
	IsAdult bool   `json:"isAdult"`
}
