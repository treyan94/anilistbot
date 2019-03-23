package anilist

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type SearchResponse struct {
	Data `json:"data"`
}

type Data struct {
	Anime `json:"anime"`
}

type Anime struct {
	Results `json:"results"`
}

type Results []Result

type Result struct {
	Description string `json:"description"`
	Title       `json:"title"`
	CoverImage  `json:"coverImage"`
	SiteUrl     string `json:"siteUrl"`
}

type Title struct {
	English       string `json:"english"`
	Native        string `json:"native"`
	Romaji        string `json:"romaji"`
	UserPreferred string `json:"userPreferred"`
}

type CoverImage struct {
	Medium     string `json:"medium"`
	Large      string `json:"large"`
	ExtraLarge string `json:"extraLarge"`
}

func Search(q string) SearchResponse {
	query := `
    query ($search: String, $isAdult: Boolean) { 
      anime: Page (perPage: 3) { 
        results: media (type: ANIME, isAdult: $isAdult, search: $search) {
          siteUrl
          title { 
           userPreferred 
          } 
          coverImage {
           medium
          }
		  description
        } 
      }
    }`

	message := map[string]interface{}{
		"query": query,
		"variables": map[string]interface{}{
			"search":  q,
			"isAdult": false,
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
