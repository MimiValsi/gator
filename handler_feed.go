package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/MimiValsi/gator/internal/database"
)

func handlerListFeeds(s *state, cmd command) error {
	ctx := context.Background()
	feeds, err := s.db.GetFeeds(ctx)
	if err != nil {
		return fmt.Errorf("couldn't list feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	fmt.Printf("Found %d feeds\n", len(feeds))
	for _, feed := range feeds {
		user, err := s.db.GetUserByID(ctx, feed.UserID)
		if err != nil {
			return fmt.Errorf("couldn't get user: %w", err)
		}
		printFeed(feed, user)
		fmt.Printf("================================")
		fmt.Println()
	}
	return nil
}

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	ctx := context.Background()

	timeNow := time.Now().UTC()
	feedName := cmd.Args[0]
	url := cmd.Args[1]
	
	createFeed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		Name:      feedName,
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(ctx, createFeed)
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		UserID: feed.UserID,
		FeedID: feed.ID,
	}
	feedFollow, err := s.db.CreateFeedFollow(ctx, feedFollowParams)
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	fmt.Println("Feed created suc2cessuflly!")
	fmt.Println("====================")
	printFeed(feed, user)
	fmt.Println("Feed following successfully:")
	printFeedFollow(feedFollow.UserName, feedFollow.FeedName)
	fmt.Println("====================")

	return nil
}

func printFeed(feed database.Feed, user database.User) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* User:          %s\n", user.Name)
	fmt.Printf("* LastFetchedAt: %v\n", feed.LastFetchedAt.Time)
}
