package app

import (
	"fmt"
)

func (c *commands) Run(s *state, cmd command) error {
	function, ok := c.cmd[cmd.name]
	if !ok {
		return fmt.Errorf("invalid command name")
	}

	err := function(s, cmd)
	if err != nil {
		return err
	}

	return nil
}

func (c *commands) Register(name string, f func(*state, command) error) error {
	_, ok := c.cmd[name]
	if ok {
		return fmt.Errorf("command already exists")
	}

	c.cmd[name] = f

	return nil
}

func CreateCommands() commands {

	c := make(map[string]func(*state, command) error)

	comms := commands{
		cmd: c,
	}

	return comms
}
