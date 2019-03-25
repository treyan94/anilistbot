package main

import (
	"anilistbot/anilist"
	"anilistbot/anilist/anime"
	"github.com/joomcode/errorx"
	"gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var apiKey = func() (key string) {
	key = os.Getenv("ANI_BOT_KEY")

	if args := os.Args[1:]; len(args) != 0 {
		key = args[0]
	}

	if key == "" {
		log.Fatal("provide telegram bot api key as first argument")
	}

	return key
}()

var bot = func() (bot *telebot.Bot) {
	bot, err := telebot.NewBot(telebot.Settings{
		Token: apiKey,
		Poller: &telebot.LongPoller{
			Timeout: 10 * time.Second,
		},
	})

	if err != nil {
		log.Fatalf("Error: %+v", errorx.Decorate(err, "this could be so much better"))
	}

	return bot
}()

func main() {
	bot.Handle("/hello", func(m *telebot.Message) {
		_, err := bot.Send(m.Sender, "hello world")

		if err != nil {
			HandleErr(err)
		}
	})

	bot.Handle(telebot.OnQuery, search)

	bot.Start()
}

func search(q *telebot.Query) {
	if q.Text == "" {
		if err := bot.Answer(q, &telebot.QueryResponse{}); err != nil {
			HandleErr(err)
		}
		return
	}

	isAdult, searchQ := parseQueryText(q.Text)

	searchResults := anilist.SearchAnime(anime.SearchVariables{
		IsAdult: isAdult,
		Search:  searchQ,
	}).Anime.Results

	err := bot.Answer(q, &telebot.QueryResponse{
		Results:   parsedResults(searchResults),
		CacheTime: 0,
	})

	if err != nil {
		HandleErr(err)
	}
}

func parsedResults(searchResults anime.Results) telebot.Results {
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

func HandleErr(err error) {
	log.Printf("Error: %+v", errorx.Decorate(err, "this could be so much better"))
}
