package anilist

type Request struct {
	Query     string `json:"query"`
	Variables `json:"variables"`
}

type Variables struct {
	Search  string `json:"search"`
	IsAdult bool   `json:"isAdult"`
}
