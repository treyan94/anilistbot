package anilist

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type SearchResponse struct {
	Data struct {
		Media `json:"media"`
	} `json:"data"`
}

type Media struct {
	Description string `json:"description"`
	Title       `json:"title"`
	CoverImage  `json:"coverImage"`
	SiteUrl     string `json:"siteUrl"`
}

type Title struct {
	English string `json:"english"`
	Native  string `json:"native"`
	Romaji  string `json:"romaji"`
}

type CoverImage struct {
	Medium     string `json:"medium"`
	Large      string `json:"large"`
	ExtraLarge string `json:"extraLarge"`
}

func Search(q string) SearchResponse {
	query := `
    query ($search: String) {
     Media (search: $search, type: ANIME) {
       title {
         romaji
         english
         native
       }
       description
       coverImage {
			extraLarge
			large
			medium
		}
		siteUrl
     }
    }`

	message := map[string]interface{}{
		"query": query,
		"variables": map[string]string{
			"search": q,
		},
	}
	jsonMarshaled, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("https://graphql.anilist.co", "application/json", bytes.NewBuffer(jsonMarshaled))
	if err != nil {
		log.Fatalln(err)
	}

	var result SearchResponse

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalln(err)
	}

	return result
}
