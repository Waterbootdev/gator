package main

import (
	"fmt"

	"github.com/Waterbootdev/gator/internal/config"
)

func main() {
	currentConfig, err := config.Read()

	if err != nil {
		fmt.Println(err)
		return
	}

	err = currentConfig.SetUser("Waterbootdev")

	if err != nil {
		fmt.Println(err)
		return
	}

	currentConfig, err = config.Read()

	if err != nil {
		fmt.Println(err)
		return
	}

	currentConfig.Print()
}
