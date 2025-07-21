package app

import (
	"context"
	"fmt"
)

func HandlerGetUsers(s *state, cmd command) error {
	currentUser := s.cfg.CurrentUserName

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		fmt.Println("error retrieving users")
		return err
	}

	for i := range users {
		if users[i] == currentUser {
			fmt.Printf("* %s (current)\n", users[i])
		} else {
			fmt.Printf("* %s\n", users[i])
		}
	}

	return nil
}
