package main

import (
	"../anilistbot/anilist"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/search", anilist.Search)
	e.Logger.Fatal(e.Start(":8000"))
}
