package character

import (
	"gopkg.in/tucnak/telebot.v2"
	"strconv"
)

type SearchResponse struct {
	Data `json:"data"`
}

type Data struct {
	Page `json:"page"`
}

type Page struct {
	Characters `json:"characters"`
}

type Characters []Character

// Returns a a formatted telegram response
func (c Characters) Parse() telebot.Results {
	parsedResults := make(telebot.Results, len(c))

	for i, character := range c {
		parsedResults[i] = &telebot.ArticleResult{
			URL:         character.SiteUrl,
			ThumbURL:    character.Image.Medium,
			Title:       character.First + " " + character.Last,
			Text:        character.SiteUrl,
			Description: character.Description,
		}
		parsedResults[i].SetResultID(strconv.Itoa(i))
	}

	return parsedResults
}

type Character struct {
	Name        `json:"name"`
	Image       `json:"image"`
	SiteUrl     string `json:"siteUrl"`
	Description string `json:"description"`
}

type Name struct {
	First string `json:"first"`
	Last  string `json:"last,omitempty"`
}

type Image struct {
	Medium string `json:"medium"`
}
