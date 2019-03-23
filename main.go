package main

import (
	"../anilistbot/anilist"
	"encoding/json"
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"strconv"
	"time"
)

type Configuration struct {
	ApiKey string `json:"API-KEY"`
}

func getConfig() (configuration Configuration) {
	file, _ := os.Open("config.json")
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

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
		err := b.Answer(q, &tb.QueryResponse{})

		if err != nil {
			fmt.Println(err)
		}
		return
	}

	media := anilist.Search(q.Text)
	results := make(tb.Results, len(media.Anime.Results))

	for i, result := range media.Anime.Results {
		res := &tb.ArticleResult{
			URL:         result.SiteUrl,
			ThumbURL:    result.CoverImage.Medium,
			Title:       result.Title.UserPreferred,
			Text:        result.SiteUrl,
			Description: result.Description,
		}
		results[i] = res
		results[i].SetResultID(strconv.Itoa(i))
	}

	err := b.Answer(q, &tb.QueryResponse{
		Results:   results,
		CacheTime: 0,
	})

	if err != nil {
		fmt.Println(err)
	}
}
