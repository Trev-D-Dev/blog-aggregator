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

func HandlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		fmt.Println("title and url of feed are required")
		os.Exit(1)
	}

	currUser, err := getUser(s)
	if err != nil {
		return err
	}

	fmt.Println("")

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    currUser.ID,
	}

	newFeed, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		fmt.Println("error creating new feed")
		return err
	}

	fmt.Println("New Feed Added:")
	fmt.Printf("ID: %v\n", newFeed.ID)
	fmt.Printf("Created At: %v\n", newFeed.CreatedAt)
	fmt.Printf("Updated At: %v\n", newFeed.UpdatedAt)
	fmt.Printf("Name: %s\n", newFeed.Name)
	fmt.Printf("URL: %s\n", newFeed.Url)
	fmt.Printf("User ID: %v\n", newFeed.UserID)

	return HandlerFollow(s, cmd)
}
