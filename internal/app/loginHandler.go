package app

import (
	"fmt"
)

func HandlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("username required for login")
	}

	username := cmd.args[0]

	err := s.cfg.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("Current user has been set to %s\n", username)

	return nil
}
