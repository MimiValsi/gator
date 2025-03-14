package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/MimiValsi/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		fmt.Printf("usage: %s <arg>", cmd.Name)
		fmt.Println("Exemple: arg 30s")
		return nil
	}
	
	timeBetweenReqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Invalid time format: %w", err)
	}
	
	fmt.Printf("Collection feeds every %v\n", timeBetweenReqs)
	ticker := time.NewTicker(timeBetweenReqs)
	for ;; <-ticker.C {
		scrapeFeeds(s)
	}
	
	//return nil
}

func scrapeFeeds(s *state) error {
	ctx := context.Background()
	feed, err := s.db.GetNextFeedtoFetch(ctx)
	if err != nil {
		return fmt.Errorf("couldn't get next feed fetch: %w", err)
	}

	tn := time.Now().UTC()
	markFeed := database.MarkFeedFetchedParams{
		UpdatedAt: tn,
		LastFetchedAt: sql.NullTime{
			Time: tn,
			Valid: true,
		},
		ID: feed.ID,
	}

	err = s.db.MarkFeedFetched(ctx, markFeed)
	if err != nil {
		return fmt.Errorf("couldn't mark fetched feed: %w", err)
	}

	rss, err := fetchRSS(ctx, feed.Url)
	if err != nil {
		return fmt.Errorf("couldn't fetch rss: %w", err)
	}

	for _, item := range rss.Channel.Items {
		fmt.Println(item.Title)
	}
	return nil
}
