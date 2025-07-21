package app

import (
	"context"
	"fmt"
	"os"
)

func HandlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("username required for login")
	}

	username := cmd.args[0]

	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		os.Exit(1)
		return fmt.Errorf("user does not exist")
	}

	err = s.cfg.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("Current user has been set to %s\n", username)

	return nil
}
