package query_parser

import "strings"

type ParsedQuery struct {
	QueryText string
	Type      string
	IsAdult   bool
}

func Parse(text string) (p ParsedQuery) {
	p.QueryText = strings.Replace(text, " /a", "", 1)
	p.QueryText = strings.Replace(p.QueryText, " /char", "", 1)

	if index := strings.Index(text, " /a"); index > -1 {
		p.IsAdult = true
	}

	p.Type = "anime"

	if index := strings.Index(text, " /char"); index > -1 {
		p.Type = "char"
	}

	return p
}
