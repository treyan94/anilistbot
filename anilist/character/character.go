package character

import (
	"anilistbot/anilist"
	"bytes"
	"encoding/json"
	"net/http"
)

func Search(searchVariables SearchVariables) (result SearchResponse) {
	query := `
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
