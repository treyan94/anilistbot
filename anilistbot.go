package main

import (
	"anilistbot/anilist"
	"gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var apiKey = func() string {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("provide telegram bot api key as first argument")
	}

	return args[0]
}()

var bot = func() (bot *telebot.Bot) {
	bot, err := telebot.NewBot(telebot.Settings{
		Token: apiKey,
		Poller: &telebot.LongPoller{
			Timeout: 10 * time.Second,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	return bot
}()

func main() {
	bot.Handle("/hello", func(m *telebot.Message) {
		_, err := bot.Send(m.Sender, "hello world")

		if err != nil {
			log.Println(err)
		}
	})

	bot.Handle(telebot.OnQuery, search)

	bot.Start()
}

func search(q *telebot.Query) {
	if q.Text == "" {
		err := bot.Answer(q, &telebot.QueryResponse{})

		if err != nil {
			log.Println(err)
		}
		return
	}

	isAdult, _ := regexp.MatchString("/a", q.Text)
	searchQ := strings.Replace(q.Text, "/a", "", 1)
	searchResults := anilist.Search(searchQ, isAdult).Anime.Results

	results := make(telebot.Results, len(searchResults))

	for i, result := range searchResults {
		res := &telebot.ArticleResult{
			URL:         result.SiteUrl,
			ThumbURL:    result.CoverImage.Medium,
			Title:       result.Title.UserPreferred,
			Text:        result.SiteUrl,
			Description: result.Description,
		}
		results[i] = res
		results[i].SetResultID(strconv.Itoa(i))
	}

	err := bot.Answer(q, &telebot.QueryResponse{
		Results:   results,
		CacheTime: 0,
	})

	if err != nil {
		log.Println(err)
	}
}
