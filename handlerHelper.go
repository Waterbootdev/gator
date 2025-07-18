package main

import (
	"context"

	"github.com/Waterbootdev/gator/internal/database"
	"github.com/google/uuid"
)

func (s *state) getUserNames() (map[uuid.UUID]string, error) {
	users, err := s.DB.GetUsers(context.Background())

	if err != nil {
		return nil, err
	}

	return userNames(users), nil
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.DB.GetUser(context.Background(), s.Config.CurrentUserName)

		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
