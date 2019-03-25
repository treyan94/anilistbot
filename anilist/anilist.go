package anilist

import (
	"anilistbot/anilist/anime"
	"bytes"
	"encoding/json"
	"github.com/joomcode/errorx"
	"log"
	"net/http"
)

func SearchAnime(searchVariables anime.SearchVariables) (result anime.SearchResponse) {
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

	reqMarshaled, err := json.Marshal(anime.Request{
		Query:           query,
		SearchVariables: searchVariables,
	})

	resp, err := http.Post("https://graphql.anilist.co", "application/json", bytes.NewBuffer(reqMarshaled))
	if err != nil {
		HandleErr(err)
	}

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		HandleErr(err)
	}

	return result
}

func HandleErr(err error) {
	log.Printf("Error: %+v", errorx.Decorate(err, "this could be so much better"))
}
