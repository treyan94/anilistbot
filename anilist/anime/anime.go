package anime

import (
	"anilistbot/anilist"
	"encoding/json"
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

func Search(searchVariables anilist.SearchVariables) (result SearchResponse) {
	cat := anilist.NewCategory(query)

	if err := json.NewDecoder(cat.RespBody(searchVariables)).Decode(&result); err != nil {
		anilist.HandleErr(err)
	}

	return result
}
