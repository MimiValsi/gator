package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	ctx := context.Background()

	err := s.db.TruncateFeeds(ctx)
	if err != nil {
		return fmt.Errorf("couln't reset feeds: %w", err)
	}
	err = s.db.TruncateUsers(ctx)
	if err != nil {
		return fmt.Errorf("couldn't reset users: %w", err)
	}

	fmt.Println("Tables successfully reset!")

	return nil
}
