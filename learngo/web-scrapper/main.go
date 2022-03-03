package main

import (
	"os"
	"strings"

	"sample/learngo/web-scrapper/scrapper"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// e.GET("/", handleHome)
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":8999"))
}

// func handleHome(c echo.Context) error {
// 	return c.String(http.StatusOK, "Hello, World!")
// }

func handleHome(c echo.Context) error {
	return c.File("public/home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove("jobs.csv")

	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment("jobs.csv", "jobs.csv")
}
