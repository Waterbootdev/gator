package main

import "github.com/Waterbootdev/gator/internal/config"

type state struct {
	Config *config.Config
}

func (s *state) setConfig() error {

	currentConfig, err := config.Read()

	if err != nil {
		return err
	}

	s.Config = &currentConfig

	return err
}
