package main

import (
	"../anilistbot/anilist"
	"encoding/json"
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"time"
)

type Configuration struct {
	ApiKey string `json:"API-KEY"`
}

func getConfig() (configuration Configuration) {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Fatal("error:", err)
	}
	return configuration
}

var b, err = tb.NewBot(tb.Settings{
	Token:  getConfig().ApiKey,
	Poller: &tb.LongPoller{Timeout: 10 * time.Second},
})

func main() {
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "hello world")
	})

	b.Handle(tb.OnQuery, search)

	b.Start()
}

func search(q *tb.Query) {
	if q.Text == "" {
		b.Answer(q, &tb.QueryResponse{})
		return
	}
	fmt.Println(q.Text)
	media := anilist.Search(q.Text).Data.Media

	res := &tb.ArticleResult{
		URL:         media.SiteUrl,
		ThumbURL:    media.CoverImage.Medium,
		Title:       media.Title.English,
		Text:        media.Description,
		Description: media.Description,
	}

	res.SetResultID("0")

	results := make(tb.Results, 1)
	results[0] = res

	err := b.Answer(q, &tb.QueryResponse{
		Results:   results,
		CacheTime: 0,
	})

	if err != nil {
		fmt.Println(err)
	}
}
