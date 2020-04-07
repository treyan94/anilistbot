package anilist

import (
	"gopkg.in/tucnak/telebot.v2"
	"log"
)

const URL = "https://graphql.anilist.co"

type Results interface {
	Parse() telebot.Results
}

type Request struct {
	Query           string `json:"query"`
	SearchVariables `json:"variables"`
}

type SearchVariables struct {
	Search  string `json:"search"`
	IsAdult bool   `json:"isAdult"`
}

func HandleErr(err error) {
	log.Printf("Error: %+v", err)
}
