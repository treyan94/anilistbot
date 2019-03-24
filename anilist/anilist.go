package anilist

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func Search(q string, isAdult bool) SearchResponse {
	query := `
    query ($search: String, $isAdult: Boolean) { 
      anime: Page (perPage: 10) { 
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

	reqMarshaled, err := json.Marshal(Request{
		Query: query,
		Variables: Variables{
			Search:  q,
			IsAdult: isAdult,
		},
	})

	resp, err := http.Post("https://graphql.anilist.co", "application/json", bytes.NewBuffer(reqMarshaled))
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
