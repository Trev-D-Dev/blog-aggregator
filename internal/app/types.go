package app

import (
	"github.com/Trev-D-Dev/blog-aggregator/internal/config"
	"github.com/Trev-D-Dev/blog-aggregator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	cmd map[string]func(*state, command) error
}
