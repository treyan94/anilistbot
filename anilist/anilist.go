package anilist

import (
	"github.com/joomcode/errorx"
	"log"
)

const URL = "https://graphql.anilist.co"

func HandleErr(err error) {
	log.Printf("Error: %+v", errorx.Decorate(err, "this could be so much better"))
}
