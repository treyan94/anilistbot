package anime

import (
	"anilistbot/anilist"
	"bytes"
	"encoding/json"
	"gopkg.in/tucnak/telebot.v2"
	"net/http"
	"strconv"
)

func ParseResults(searchResults Results) telebot.Results {
	parsedResults := make(telebot.Results, len(searchResults))

	for i, result := range searchResults {
		parsedResults[i] = &telebot.ArticleResult{
			URL:         result.SiteUrl,
			ThumbURL:    result.CoverImage.Medium,
			Title:       result.Title.UserPreferred,
			Text:        result.SiteUrl,
			Description: result.Description,
		}
		parsedResults[i].SetResultID(strconv.Itoa(i))
	}

	return parsedResults
}

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
