package app

import (
	"context"
	"fmt"

	"github.com/Trev-D-Dev/blog-aggregator/internal/database"
)

func HandlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("must enter a url to unfollow")
	}

	params := database.DeleteFeedFollowParams{
		UserID: user.ID,
		Url:    cmd.args[0],
	}

	err := s.db.DeleteFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error unfollowing feed: %v", err)
	}

	fmt.Printf("successfully unfollowed '%s'\n", cmd.args[0])

	return nil
}
