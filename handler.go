package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Waterbootdev/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {

	if len(cmd.arguments) == 0 {
		return errors.New("please provide a username")
	}

	user, err := s.DB.GetUser(context.Background(), cmd.arguments[0])

	if err != nil {
		fmt.Println("no user found")
		os.Exit(1)
	}

	if user.Name != cmd.arguments[0] {
		fmt.Println("Fatal error occured while logging in")
		os.Exit(1)
	}

	err = s.Config.SetUser(cmd.arguments[0])

	if err != nil {
		return err
	}

	fmt.Println("Login successful")

	return nil
}

func handleRegister(s *state, cmd command) error {

	if len(cmd.arguments) == 0 {
		return errors.New("please provide a username")
	}

	currentTime := time.Now()

	currentUser, err := s.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      cmd.arguments[0],
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return s.Config.SetUser(currentUser.Name)
}

func handleReset(s *state, _ command) error {
	return s.DB.DeleteALLUsers(context.Background())
}

func handleGetUsers(s *state, _ command) error {

	users, err := s.DB.GetUsers(context.Background())

	if err != nil {
		return err
	}

	curentUserName := s.Config.CurrentUserName

	for _, user := range users {
		if user.Name == curentUserName {
			fmt.Println(user.Name, "(current)")
		} else {
			fmt.Println(user.Name)
		}
	}

	return nil
}

func (c *commands) registerAll() {

	c.availableCommands = map[string]func(s *state, cmd command) error{}
	c.register("login", handlerLogin)
	c.register("register", handleRegister)
	c.register("reset", handleReset)
	c.register("users", handleGetUsers)
}
