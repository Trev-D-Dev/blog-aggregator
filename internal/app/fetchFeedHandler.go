package app

import (
	"fmt"
	"time"
)

func HandlerFetchFeed(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("must include a time between requests")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("error parsing time duration: %v", err)
	}

	fmt.Printf("Collecting feeds every %s\n", cmd.args[0])

	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

	return nil
}
