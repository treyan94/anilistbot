package main

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
			Description string
			Title       struct {
				English string
				Native  string
				Romaji  string
			}
		}
	}
}

func main() {
	e := echo.New()

	e.GET("/search", search)
	e.Logger.Fatal(e.Start(":8000"))
}

func search(c echo.Context) error {
	query := `
    query ($search: String) {
     Media (search: $search, type: ANIME) {
       title {
         romaji
         english
         native
       }
       description
     }
    }`

	message := map[string]interface{}{
		"query": query,
		"variables": map[string]string{
			"search": "Evangelion",
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

	json.NewDecoder(resp.Body).Decode(&result)

	return c.JSON(http.StatusOK, result)
}
