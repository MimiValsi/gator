package main

import (
	"context"
	"fmt"
	"strconv"
)

func handlerBrowse(s *state, cmd command) error {
	limit := 0
	var err error
	if len(cmd.Args) < 1 {
		fmt.Printf("usage: %s limit <optional>\n", cmd.Name)
		fmt.Println("Default limit = 2")
	}

	if len(cmd.Args) == 1 {
		limit = 2
	} else {
		limit, err = strconv.Atoi(cmd.Args[1])
		if err != nil {
			return fmt.Errorf("wrong number format: %w", err)
		}
	}
	fmt.Println("limit:", limit)
	posts, err := s.db.GetPostsPerUser(context.Background(), int32(limit))
	if err != nil {
		return fmt.Errorf("couldn't get posts: %w", err)
	}

	for _, post := range posts {
		fmt.Printf("> %s\n", post.Description)
	}
	return nil
}
