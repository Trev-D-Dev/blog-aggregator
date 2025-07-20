package main

import (
	"fmt"
	"os"

	"github.com/Trev-D-Dev/blog-aggregator/internal/app"
	"github.com/Trev-D-Dev/blog-aggregator/internal/config"
)

func main() {
	config, err := config.Read()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	state := app.CreateState(&config)

	comms := app.CreateCommands()

	comms.Register("login", app.HandlerLogin)

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
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
