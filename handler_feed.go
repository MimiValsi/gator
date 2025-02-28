package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/MimiValsi/gator/internal/database"
)

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
