package character

import (
	"anilistbot/anilist"
	"encoding/json"
)

const query string = `
    query ($search: String) {
      page: Page(perPage: 10) {
        characters: characters(search: $search) {
          name {
            first
            last
          }
          image {
            medium
          }
          siteUrl
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
