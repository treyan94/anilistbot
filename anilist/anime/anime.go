package anime

import (
	"anilistbot/anilist"
	"bytes"
	"encoding/json"
	"net/http"
)

const query string = `
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

func Search(searchVariables SearchVariables) (result SearchResponse) {

	reqMarshaled, err := json.Marshal(Request{
		Query:           query,
		SearchVariables: searchVariables,
	})

	resp, err := http.Post(anilist.URL, "application/json", bytes.NewBuffer(reqMarshaled))
	if err != nil {
		anilist.HandleErr(err)
	}

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		anilist.HandleErr(err)
	}

	return result
}
