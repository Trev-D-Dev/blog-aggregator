package app

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Trev-D-Dev/blog-aggregator/internal/rss"
)

func HandlerFetchFeed(s *state, cmd command) error {

	url := "https://www.wagslane.dev/index.xml"

	rssFeed, err := rss.FetchFeed(context.Background(), url)
	if err != nil {
		fmt.Println("error fetching feed")
		return err
	}

	b, err := json.MarshalIndent(rssFeed, "", " ")
	if err != nil {
		fmt.Println("error with json.MarshalIndent")
		return err
	}

	fmt.Println(string(b))

	return nil
}
