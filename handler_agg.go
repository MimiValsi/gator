package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	ctx := context.Background()
	url := "https://www.wagslane.dev/index.xml"
	rssFeed, err := fetchRSS(ctx, url)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", rssFeed)
	return nil
}
