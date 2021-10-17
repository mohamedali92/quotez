package main

import "time"

// Quote represents a quote object
type Quote struct {
	Id        int
	CreatedAt time.Time
	QuoteText string
	Author    string
	Tags      []string
	Likes     int
	QuoteUrl  string
}
