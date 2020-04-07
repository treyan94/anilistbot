package query_parser

import "strings"

const adult string = " /a"
const character string = " /char"

type ParsedQuery struct {
	QueryText string
	Type      string
	IsAdult   bool
}

func Parse(text string) (p ParsedQuery) {
	p.QueryText = strings.Replace(text, adult, "", 1)
	p.QueryText = strings.Replace(p.QueryText, character, "", 1)

	if strings.Contains(text, adult) {
		p.IsAdult = true
	}

	p.Type = "anime"

	if strings.Contains(text, character) {
		p.Type = "char"
	}

	return p
}
