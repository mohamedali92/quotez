package main

import (
	"context"
	"github.com/jackc/pgx/v4"
)

func insertQuote(ctx context.Context, conn *pgx.Conn, quote Quote) error {
	sql := `INSERT INTO public.quotes 
			(id, created_at, quote_text, author, tags, likes, quote_url) 
			VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := conn.Exec(ctx, sql, quote.Id, quote.CreatedAt, quote.QuoteText, quote.Author, quote.Tags, quote.Likes, quote.QuoteUrl)
	if err != nil {
		return err
	} else {
		return nil
	}
}
