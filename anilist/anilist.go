package anilist

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func Search(searchVariables SearchVariables) (result SearchResponse) {
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
		query,
		searchVariables,
	})

	resp, err := http.Post("https://graphql.anilist.co", "application/json", bytes.NewBuffer(reqMarshaled))
	if err != nil {
		log.Fatalln(err)
	}

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	return result
}
