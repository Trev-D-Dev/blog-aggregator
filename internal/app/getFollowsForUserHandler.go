package app

import (
	"context"
	"fmt"

	"github.com/Trev-D-Dev/blog-aggregator/internal/database"
)

func HandlerGetFeedFollowsForUser(s *state, cmd command, user database.User) error {

	uID := user.ID

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), uID)
	if err != nil {
		return fmt.Errorf("error retrieving follows for user '%s': %v", s.cfg.CurrentUserName, err)
	}

	fmt.Printf("User '%s' is following: \n", s.cfg.CurrentUserName)

	for i := range follows {
		fmt.Printf(" - %s\n", follows[i].FeedName)
	}

	return nil
}
