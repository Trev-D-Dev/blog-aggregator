package app

import (
	"github.com/Trev-D-Dev/blog-aggregator/internal/config"
	"github.com/Trev-D-Dev/blog-aggregator/internal/database"
)

func CreateState(c *config.Config, db *database.Queries) state {
	s := state{
		db:  db,
		cfg: c,
	}

	return s
}
