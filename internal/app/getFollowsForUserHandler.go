package app

import (
	"context"
	"fmt"
)

func HandlerGetFeedFollowsForUser(s *state, cmd command) error {

	currUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error retrieving user: %v", err)
	}

	uID := currUser.ID

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
