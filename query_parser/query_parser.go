package query_parser

import "strings"
import "anilistbot/anilist"
import "anilistbot/anilist/anime"
import "anilistbot/anilist/character"

const adult string = " /a"
const char string = " /char"

type ParsedQuery struct {
	QueryText string
	Type      string
	IsAdult   bool
}

func parse(text string) (p ParsedQuery) {
	p.QueryText = strings.Replace(text, adult, "", 1)
	p.QueryText = strings.Replace(p.QueryText, char, "", 1)

	if strings.Contains(text, adult) {
		p.IsAdult = true
	}

	p.Type = "anime"

	if strings.Contains(text, char) {
		p.Type = "char"
	}

	return p
}

func Parse(text string) (r anilist.Results) {
	p := parse(text)

	sv := anilist.SearchVariables{
		IsAdult: p.IsAdult,
		Search:  p.QueryText,
	}

	if p.Type == "anime" {
		return anime.Search(sv)
	}

	return character.Search(sv)
}
