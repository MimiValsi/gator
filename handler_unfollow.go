package main

import (
	"context"
	"fmt"
	
	"github.com/MimiValsi/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.Name)
	}

	ctx := context.Background()
	feed, err := s.db.GetFeedByURL(ctx, cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}
	
	del := database.DeleteFeedByUserAndURLParams{
		FeedID: feed.ID,
		UserID: user.ID,
	}
	
	err = s.db.DeleteFeedByUserAndURL(ctx, del)
	if err != nil {
		return fmt.Errorf("couldn't delete feed follow: %w", err)
	}

	return nil
}
