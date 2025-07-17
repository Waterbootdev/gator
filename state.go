package main

import (
	"database/sql"

	"github.com/Waterbootdev/gator/internal/config"
	"github.com/Waterbootdev/gator/internal/database"
)

type state struct {
	DB     *database.Queries
	Config *config.Config
}

func (s *state) setDBQueries() error {

	db, err := sql.Open("postgres", s.Config.DBUrl)

	if err != nil {
		return err
	}

	s.DB = database.New(db)

	return err
}

func (s *state) setConfig() error {

	currentConfig, err := config.Read()

	if err != nil {
		return err
	}

	s.Config = &currentConfig

	return err
}
