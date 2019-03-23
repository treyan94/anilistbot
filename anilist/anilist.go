package anilist

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

type searchResponse struct {
	Data struct {
		Media struct {
			Description string `json:"description"`
			Title       struct {
				English string `json:"english"`
				Native  string `json:"native"`
				Romaji  string `json:"romaji"`
			} `json:"title"`
			CoverImage struct {
				Large      string `json:"large"`
				ExtraLarge string `json:"extraLarge"`
			} `json:"coverImage"`
		} `json:"media"`
	} `json:"data"`
}

func Search(c echo.Context) error {
	query := `
    query ($search: String) {
     Media (search: $search, type: ANIME) {
       title {
         romaji
         english
         native
       }
       description
       coverImage {
			extraLarge
			large
		}
     }
    }`

	message := map[string]interface{}{
		"query": query,
		"variables": map[string]string{
			"search": c.QueryParam("q"),
		},
	}
	jsonMarshaled, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("https://graphql.anilist.co", "application/json", bytes.NewBuffer(jsonMarshaled))
	if err != nil {
		log.Fatalln(err)
	}

	var result searchResponse

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalln(err)
	}

	return c.JSON(http.StatusOK, result)
}
