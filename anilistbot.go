package main

import (
	"anilistbot/anilist"
	"anilistbot/anilist/anime"
	"anilistbot/anilist/character"
	"anilistbot/query_parser"
	"gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
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
		log.Fatalf("Error: %+v", err)
	}

	return bot
}()

func main() {
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

	parsedQuery := query_parser.Parse(q.Text)
	searchResults := *new(anilist.Results)

	switch parsedQuery.Type {
	case "anime":
		searchResults = anime.Search(anime.SearchVariables{
			IsAdult: parsedQuery.IsAdult,
			Search:  parsedQuery.QueryText,
		})
	case "char":
		searchResults = character.Search(character.SearchVariables{
			Search: parsedQuery.QueryText,
		})
	}

	parsed := searchResults.Parse()

	err := bot.Answer(q, &telebot.QueryResponse{
		Results:   parsed,
		CacheTime: 0,
	})

	if err != nil {
		HandleErr(err)
	}
}

func HandleErr(err error) {
	log.Printf("Error: %+v", err)
}
