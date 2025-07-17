package main

import (
	"fmt"
	"os"

	"github.com/Waterbootdev/gator/internal/config"
)

func initialize() (*state, *commands) {

	currentConfig, err := config.Read()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	currentState := &state{
		Config: &currentConfig,
	}

	commands := &commands{
		availableCommands: map[string]func(s *state, cmd command) error{},
	}

	commands.register("login", handlerLogin)

	return currentState, commands
}

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

func main() {

	currentState, commands := initialize()

	err := commands.run(currentState, getCommand())

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	currentState.Config.Print()
}
