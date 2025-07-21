package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Trev-D-Dev/blog-aggregator/internal/app"
	"github.com/Trev-D-Dev/blog-aggregator/internal/config"
	"github.com/Trev-D-Dev/blog-aggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.Read()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", config.URL)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	dbQueries := database.New(db)

	state := app.CreateState(&config, dbQueries)

	comms := app.CreateCommands()

	comms.Register("login", app.HandlerLogin)
	comms.Register("register", app.HandlerRegister)
	comms.Register("reset", app.HandlerReset)
	comms.Register("users", app.HandlerGetUsers)
	comms.Register("agg", app.HandlerFetchFeed)
	comms.Register("addfeed", app.MiddlewareLoggedIn(app.HandlerAddFeed))
	comms.Register("feeds", app.HandlerGetFeeds)
	comms.Register("follow", app.MiddlewareLoggedIn(app.HandlerFollow))
	comms.Register("following", app.MiddlewareLoggedIn(app.HandlerGetFeedFollowsForUser))

	args := os.Args
	if len(args) < 2 {
		fmt.Println("requires at least 2 commands")
		os.Exit(1)
	}

	cmdName := args[1]
	cmdArgs := args[2:]
	newCmd := app.CreateCommand(cmdName, cmdArgs)

	err = comms.Run(&state, newCmd)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
