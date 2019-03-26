package anilist

import (
	"github.com/joomcode/errorx"
	"gopkg.in/tucnak/telebot.v2"
	"log"
)

type Results interface {
	Parse() telebot.Results
}

const URL = "https://graphql.anilist.co"

func HandleErr(err error) {
	log.Printf("Error: %+v", errorx.Decorate(err, "this could be so much better"))
}
