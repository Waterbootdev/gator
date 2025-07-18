package main

import (
	"github.com/Waterbootdev/gator/internal/database"
	"github.com/google/uuid"
)

func userNames(users []database.User) map[uuid.UUID]string {
	m := make(map[uuid.UUID]string, len(users))

	for _, user := range users {
		m[user.ID] = user.Name
	}

	return m
}
