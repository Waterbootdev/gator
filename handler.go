package main

import (
	"errors"
	"fmt"
)

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

func (c *commands) registerAll() {

	c.availableCommands = map[string]func(s *state, cmd command) error{}
	c.register("login", handlerLogin)
}
