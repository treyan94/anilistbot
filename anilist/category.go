package anilist

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Category struct {
	Query string
}

func NewCategory(query string) *Category {
	return &Category{Query: query}
}

// Performs the search request and returns the response body
func (c *Category) RespBody(searchVariables SearchVariables) io.ReadCloser {
	reqMarshaled, err := json.Marshal(Request{
		Query:           c.Query,
		SearchVariables: searchVariables,
	})

	resp, err := http.Post(URL, "application/json", bytes.NewBuffer(reqMarshaled))
	if err != nil {
		HandleErr(err)
	}

	return resp.Body
}
