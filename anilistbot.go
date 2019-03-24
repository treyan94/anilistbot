package main

import (
	"anilistbot/anilist"
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var bot = getBot()
var apiKey = getApiKey()

func getBot() (bot *tb.Bot) {
	bot, err := tb.NewBot(tb.Settings{
		Token: apiKey,
		Poller: &tb.LongPoller{
			Timeout: 10 * time.Second,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	return bot
}

func getApiKey() string {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("provide telegram bot api key as first argument")
	}

	return args[0]
}

func main() {
	bot.Handle("/hello", func(m *tb.Message) {
		_, err := bot.Send(m.Sender, "hello world")

		if err != nil {
			fmt.Println(err)
		}
	})

	bot.Handle(tb.OnQuery, search)

	bot.Start()
}

func search(q *tb.Query) {
	if q.Text == "" {
		err := bot.Answer(q, &tb.QueryResponse{})

		if err != nil {
			fmt.Println(err)
		}
		return
	}

	isAdult, _ := regexp.MatchString("/a", q.Text)
	searchQ := strings.Replace(q.Text, "/a", "", 1)
	media := anilist.Search(searchQ, isAdult)

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

	err := bot.Answer(q, &tb.QueryResponse{
		Results:   results,
		CacheTime: 0,
	})

	if err != nil {
		fmt.Println(err)
	}
}
