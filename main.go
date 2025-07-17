package main

import (
	"fmt"
	"os"
)

func main() {

	currentState, currentCommands := initializeOrExit()

	err := currentCommands.run(&currentState, getCommand())

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	currentState.Config.Print()
}
