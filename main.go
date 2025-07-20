package main

import (
	"fmt"

	"github.com/Trev-D-Dev/blog-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	cfg.SetUser("trevor")

	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	fmt.Println("Contents:")

	fmt.Printf("   URL: %s\n", cfg.URL)
	fmt.Printf("   User: %s\n", cfg.CurrentUserName)
}
