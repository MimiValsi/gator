package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	ctx := context.Background()

	err := s.db.DeleteFeedFollows(ctx)
	if err != nil {
		return fmt.Errorf("couldn't reset table: %w", err)
	}
	err = s.db.DeleteUsers(ctx)
	if err != nil {
		return fmt.Errorf("couldn't reset table: %w", err)
	}

	fmt.Println("Table successfully reset!")

	return nil
}
