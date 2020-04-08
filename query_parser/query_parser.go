package query_parser

import "strings"
import "anilistbot/anilist"
import "anilistbot/anilist/anime"
import "anilistbot/anilist/character"

const adult string = " /a"
const char string = " /char"

type ParsedQuery struct {
	QueryText string // Search query text
	Type      string // Search type (currently only anime or character
	IsAdult   bool   // Indicates of the search should include adult content
}

func (p ParsedQuery) ToSearchVariable() anilist.SearchVariables {
	return anilist.SearchVariables{
		IsAdult: p.IsAdult,
		Search:  p.QueryText,
	}
}

func parse(text string) (p ParsedQuery) {
	p.QueryText = strings.Replace(text, adult, "", 1)
	p.QueryText = strings.Replace(p.QueryText, char, "", 1)

	if strings.Contains(text, adult) {
		p.IsAdult = true
	}

	p.Type = "anime" // Search type default to anime

	if strings.Contains(text, char) {
		p.Type = "char"
	}

	return p
}

func Parse(text string) (r anilist.Results) {
	p := parse(text)

	sv := p.ToSearchVariable()

	if p.Type == "anime" {
		return anime.Search(sv)
	}

	return character.Search(sv)
}
