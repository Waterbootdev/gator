package main

import (
	"fmt"
)

type command struct {
	name      string
	arguments []string
}

type commands struct {
	availableCommands map[string]func(s *state, cmd command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.availableCommands[cmd.name]

	if !ok {
		return fmt.Errorf("command %s not found", cmd.name)
	}

	return handler(s, cmd)
}

func (c *commands) register(name string, handler func(s *state, cmd command) error) {
	c.availableCommands[name] = handler
}
