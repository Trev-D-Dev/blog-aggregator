package app

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Trev-D-Dev/blog-aggregator/internal/database"
)

func HandlerBrowse(s *state, cmd command, user database.User) error {

	var limit int
	var err error

	if len(cmd.args) < 1 {
		limit = 2
	} else {
		limit, err = strconv.Atoi(cmd.args[0])
		if err != nil {
			fmt.Println("unable to convert string to int, setting limit to default")
			limit = 2
		}
	}

	params := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}

	posts, err := s.db.GetPostsForUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error retrieving posts: %v", err)
	}

	for i := range posts {
		fmt.Printf("Title: %s\n", posts[i].Title.String)
		fmt.Printf("Published: %v\n", posts[i].PublishedAt.Time)
		fmt.Printf("Description: %s\n", posts[i].Description.String)
		fmt.Printf("URL: %s\n", posts[i].Url)
		fmt.Println("")
	}

	return nil
}
