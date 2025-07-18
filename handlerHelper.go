package main

import (
	"context"

	"github.com/google/uuid"
)

func (s *state) getUserNames() (map[uuid.UUID]string, error) {
	users, err := s.DB.GetUsers(context.Background())

	if err != nil {
		return nil, err
	}

	return userNames(users), nil
}
