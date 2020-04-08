package main

import (
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
		log.Fatal("Error connecting to telegram, please check your api key.")
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

	err := bot.Answer(q, &telebot.QueryResponse{
		Results:   query_parser.Parse(q.Text).Parse(),
		CacheTime: 0,
	})

	if err != nil {
		HandleErr(err)
	}
}

func HandleErr(err error) {
	log.Printf("Error: %+v", err)
}
