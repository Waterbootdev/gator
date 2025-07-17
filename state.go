package main

import (
	"errors"
	"fmt"

	"github.com/Waterbootdev/gator/internal/config"
)

type state struct {
	Config *config.Config
}

type command struct {
	name      string
	arguments []string
}

type commands struct {
	availableCommands map[string]func(s *state, cmd command) error
}

func handlerLogin(s *state, cmd command) error {

	if len(cmd.arguments) == 0 {
		return errors.New("please provide a username")
	}

	err := s.Config.SetUser(cmd.arguments[0])

	if err != nil {
		return err
	}

	fmt.Println("Login successful")
	return nil

}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.availableCommands[cmd.name]

	if !ok {
		return nil
	}

	return handler(s, cmd)
}

func (c *commands) register(name string, handler func(s *state, cmd command) error) {
	c.availableCommands[name] = handler
}
