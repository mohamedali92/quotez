package main

import (
	"context"
	"github.com/jackc/pgx/v4"

	//"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var BaseUrl = "https://www.goodreads.com"

func main() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//log.SetOutput(file)

	// setup up db connection
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, os.Getenv("QUOTESDSN"))
	Check(err)
	defer conn.Close(ctx)

	var quotes []Quote
	c := colly.NewCollector(colly.AllowedDomains("www.goodreads.com"))

	c.OnHTML(".quoteDetails", func(e *colly.HTMLElement) {
		quoteText := processQuoteContent(e.ChildText("div.quoteText"))
		author := cleanAuthorField(strings.TrimSpace(e.ChildText(".authorOrTitle")))
		likes, err := strconv.Atoi(strings.TrimSuffix(e.ChildText(".right"), " likes"))
		if err != nil {
			log.Println("Failed to extract number of likes ", err)
		}
		tags := extractTags(e.ChildText("div.greyText"))
		quoteUrl := BaseUrl + e.ChildAttr("a.smallText", "href")
		id := extractId(quoteUrl)
		q := Quote{Id: id, CreatedAt: time.Now(), QuoteText: quoteText, Author: author, Tags: tags, Likes: likes, QuoteUrl: quoteUrl}
		quotes = append(quotes, q)
	})

	log.Print("Logging to a file in Go!")

	log.Println("Starting Scraper")
	c.Visit(BaseUrl + "/quotes/")

	log.Println(quotes, len(quotes))

}
