package config

import (
	"encoding/json"
	"os"
)

func getConfigFileName() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return home + "/.gatorconfig.json", nil
}

func (c *Config) write() error {

	configFileName, err := getConfigFileName()

	if err != nil {
		return err
	}

	file, err := os.Create(configFileName)

	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	return encoder.Encode(c)
}
