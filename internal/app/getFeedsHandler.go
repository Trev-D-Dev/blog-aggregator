package app

import (
	"context"
	"fmt"
)

func HandlerGetFeeds(s *state, cmd command) error {

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		fmt.Println("error retrieving feeds")
	}

	fmt.Println("Feeds:")

	for i := range feeds {
		user, err := s.db.GetUserFromID(context.Background(), feeds[i].UserID)
		if err != nil {
			fmt.Printf("error retrieving user from id '%v'\n", feeds[i].UserID)
			return err
		}

		fmt.Printf(" - Name: %s\n", feeds[i].Name)
		fmt.Printf("   URL: %s\n", feeds[i].Url)
		fmt.Printf("   User: %s\n", user.Name)
	}

	return nil
}
