package app

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Trev-D-Dev/blog-aggregator/internal/database"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func HandlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("must enter a name")
	}

	// Checks if user with passed name already exists
	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err == nil {
		fmt.Printf("User with name '%s' already exists\n", cmd.args[0])
		os.Exit(1)
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}

	newUser, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		return err
	}

	s.cfg.SetUser(newUser.Name)

	fmt.Println("New User Created:")
	fmt.Printf("Name: %s\n", newUser.Name)
	fmt.Printf("CreatedAt: %v\n", newUser.CreatedAt)
	fmt.Printf("UpdatedAt: %v\n", newUser.UpdatedAt)
	fmt.Printf("ID: %v\n", newUser.ID)

	return nil
}
