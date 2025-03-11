package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/MimiValsi/gator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name> <feed_url>", cmd.Name)
	}

	// cmd.Args[0] => url
	ctx := context.Background()
	feed, err := s.db.GetFeedByURL(ctx, cmd.Args[0])
	if err != nil {
		return fmt.Errorf("coulnd't get feed: %w", err)
	}

	user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("can't fetch user: %w", err)
	}
	
	timeNow := time.Now().UTC()
	feedParams := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		UserID: user.ID,
		FeedID: feed.ID,
	}
	
	//fmt.Printf("feed name: %s\n", feed)
	//fmt.Printf("feed name: %+v\n", feedParams)
	ffRow, err = s.db.CreateFeedFollow(ctx, feedParams)
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}


	fmt.Println("Feed follow created:")
	printFeedFollow(ffRow.UserName, ffRow.FeedName)

	return nil
}

func handlerListFeedFollow(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feed follow: %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}

	fmt.Printf("Feed follows for user %s:", user.Name)
	for _, ff := range feedFollows {
		fmt.Printf("* %s\n", ff.FeedName)
	}

	return nil
}

func printFeedFollow(username, feedname string) {
	fmt.Printf("* User:          %s\n", username)
	fmt.Printf("* Feed:          %s\n", feedname)
}
