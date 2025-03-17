package main

import (
	"context"
	"fmt"
	"time"
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
	
	fmt.Printf("Collection feeds every %v...\n", timeBetweenReqs)
	ticker := time.NewTicker(timeBetweenReqs)
	for ;; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	ctx := context.Background()
	feed, err := s.db.GetNextFeedtoFetch(ctx)
	if err != nil {
		return fmt.Errorf("couldn't get next feed fetch: %w", err)
	}

	_, err = s.db.MarkFeedFetched(ctx, feed.ID)
	if err != nil {
		return fmt.Errorf("couldn't mark fetched feed: %w", err)
	}

	rss, err := fetchRSS(ctx, feed.Url)
	if err != nil {
		return fmt.Errorf("couldn't fetch rss: %w", err)
	}

	for _, item := range rss.Channel.Item {
		fmt.Println(item.Title)
	}
	fmt.Printf("Feed %s collected, %v posts found\n", feed.Name, len(rss.Channel.Item))
	return nil
}
