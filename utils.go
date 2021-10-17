package main

import (
	"log"
	"strconv"
	"strings"
	"unicode/utf8"
)

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// processQuoteContent parses the text blob to extract the quote text and author
func processQuoteContent(content string) string {
	const cdata = "//<![CDATA"
	if strings.Contains(content, cdata) {
		index := strings.Index(content, cdata)
		content = strings.TrimSpace(content[:index])
	}
	contents := strings.Split(content, "  ―\n  ")

	quoteText := strings.TrimSpace(contents[0])
	lastIndex := utf8.RuneCountInString(quoteText)
	quoteText = quoteText[:lastIndex]
	quoteText = strings.TrimPrefix(quoteText, "“")
	return quoteText
}

func cleanAuthorField(author string) string {

	index := strings.Index(author, ",")

	if index == -1 {
		return author
	} else {
		return author[:index]
	}
}

func extractTags(tagsString string) []string {
	trimmed := strings.TrimPrefix(tagsString, "tags:\n")
	trimmed = strings.ReplaceAll(trimmed, "       ", "")
	tags := strings.Split(trimmed, ", ")
	return tags

}

func extractId(quoteUrl string) int {
	trimmed := strings.TrimPrefix(quoteUrl, "https://www.goodreads.com/quotes/")
	id := strings.SplitN(trimmed, "-", 2)[0]
	id_int, err := strconv.Atoi(id)

	if err != nil {
		log.Println("Failed to extract id from quoteUrl: ", quoteUrl, err)
	}
	return id_int
}
