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

	isAdult, searchQ := parseQueryText(q.Text)
	searchResults := anilist.Search(searchQ, isAdult).Anime.Results

	err := bot.Answer(q, &telebot.QueryResponse{
		Results:   parsedResults(searchResults),
		CacheTime: 0,
	})

	if err != nil {
		log.Println(err)
	}
}

func parsedResults(searchResults anilist.Results) telebot.Results {
	parsedResults := make(telebot.Results, len(searchResults))

	for i, result := range searchResults {
		parsedResults[i] = &telebot.ArticleResult{
			URL:         result.SiteUrl,
			ThumbURL:    result.CoverImage.Medium,
			Title:       result.Title.UserPreferred,
			Text:        result.SiteUrl,
			Description: result.Description,
		}
		parsedResults[i].SetResultID(strconv.Itoa(i))
	}

	return parsedResults
}

func parseQueryText(text string) (isAdult bool, query string) {
	isAdult, _ = regexp.MatchString("/a", text)
	query = strings.Replace(text, "/a", "", 1)

	return isAdult, query
}
