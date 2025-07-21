package app

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Trev-D-Dev/blog-aggregator/internal/database"
	"github.com/Trev-D-Dev/blog-aggregator/internal/rss"
	"github.com/araddon/dateparse"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func scrapeFeeds(s *state) error {

	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error retrieving next feed: %v", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %v", err)
	}

	rssFeed, err := rss.FetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		fmt.Println("error fetching feed")
		return err
	}

	for i := range rssFeed.Channel.Item {

		title := rssFeed.Channel.Item[i].Title
		nsTitle := sql.NullString{String: title, Valid: true}
		url := rssFeed.Channel.Item[i].Link
		desc := rssFeed.Channel.Item[i].Description
		nsDesc := sql.NullString{String: desc, Valid: true}
		pubDate := rssFeed.Channel.Item[i].PubDate
		parTime, err := dateparse.ParseAny(pubDate)
		if err != nil {
			fmt.Println("error parsing publish date")
			parTime = time.Now()
		}
		ntParTime := sql.NullTime{Time: parTime, Valid: true}

		params := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       nsTitle,
			Url:         url,
			Description: nsDesc,
			PublishedAt: ntParTime,
			FeedID:      nextFeed.ID,
		}

		_, err = s.db.CreatePost(context.Background(), params)
		if err != nil {
			return fmt.Errorf("error creating post: %v", err)
		}
	}

	return nil
}
