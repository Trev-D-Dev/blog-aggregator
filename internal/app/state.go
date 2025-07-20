package app

import (
	"github.com/Trev-D-Dev/blog-aggregator/internal/config"
)

func CreateState(c *config.Config) state {
	s := state{
		cfg: c,
	}

	return s
}
