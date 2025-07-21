package app

import (
	"context"
	"fmt"

	"github.com/Trev-D-Dev/blog-aggregator/internal/database"
)

func getUser(s *state) (database.User, error) {
	currUserName := s.cfg.CurrentUserName

	currUser, err := s.db.GetUser(context.Background(), currUserName)
	if err != nil {
		fmt.Println("error retrieving user '%s'", currUserName)
		return database.User{}, err
	}

	return currUser, nil
}

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
