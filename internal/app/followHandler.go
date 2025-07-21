package app

import (
	"context"
	"fmt"

	"github.com/Trev-D-Dev/blog-aggregator/internal/database"
)

func HandlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("must include only one url to follow")
	}

	argsLen := len(cmd.args)

	feedURL := cmd.args[argsLen-1]

	feed, err := s.db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("error retrieving feed: %v", err)
	}

	params := database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error creating feedFollow: %v", err)
	}

	fmt.Printf("Feed: %s, User: %s\n", feedFollow.FeedName, feedFollow.UserName)

	return nil
}
