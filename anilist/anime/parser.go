package anime

import (
	"gopkg.in/tucnak/telebot.v2"
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
