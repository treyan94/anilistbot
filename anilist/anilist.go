package anilist

import (
	"gopkg.in/tucnak/telebot.v2"
	"log"
)

const URL = "https://graphql.anilist.co"

type Results interface {
	Parse() telebot.Results
}

func HandleErr(err error) {
	log.Printf("Error: %+v", err)
}
