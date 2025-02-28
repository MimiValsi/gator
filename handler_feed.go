package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/MimiValsi/gator/internal/database"
)

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't list feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("Feed Name: %v\n", feed.Name)
		fmt.Printf("Feed Url:  %v\n", feed.Url)
		fmt.Printf("User Name: %v\n", feed.Name_2)
		fmt.Println()
	}
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	// s.cfg.CurrentUserName
	ctx := context.Background()
	current, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	timeNow := time.Now().UTC()
	feedName := cmd.Args[0]
	url := cmd.Args[1]
	createFeed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		Name:      feedName,
		Url:       url,
		UserID:    current.ID,
	}

	feed, err := s.db.CreateFeed(ctx, createFeed)
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	fmt.Println("Feed created successuflly!")
	fmt.Println("====================")
	printFeed(feed)
	fmt.Println("====================")

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}
