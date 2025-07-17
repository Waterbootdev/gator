package main

import (
	"fmt"
	"os"
)

func getCommand() command {
	args := os.Args

	if len(args) < 2 {
		fmt.Print("pleace provide a command")
		os.Exit(1)
	}

	cmd := command{
		name:      args[1],
		arguments: args[2:],
	}

	return cmd
}

func initializeOrExit() (state, commands) {
	currentState := state{}
	err := currentState.setConfig()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	currentCommands := commands{}

	currentCommands.registerAll()

	return currentState, currentCommands
}
