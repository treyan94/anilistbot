package anime

import (
	"gopkg.in/tucnak/telebot.v2"
	"strconv"
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

func (r Results) Parse() telebot.Results {
	parsedResults := make(telebot.Results, len(r))

	for i, result := range r {
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
