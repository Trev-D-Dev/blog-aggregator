package app

import (
	"context"
	"fmt"
	"os"
)

func HandlerReset(s *state, cmd command) error {

	err := s.db.ResetUsers(context.Background())
	if err != nil {
		fmt.Println("users failed to reset")
		os.Exit(1)
	}

	fmt.Println("users successfully reset")
	os.Exit(0)
	return nil
}
