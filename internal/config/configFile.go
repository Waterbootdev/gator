package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName string = ".gatorconfig.json"

func getConfigFileName() (string, error) {

	home, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	return filepath.Join(home, configFileName), err
}

func (c *Config) encode(configFileName string) error {

	file, err := os.Create(configFileName)

	if err != nil {
		return err
	}

	defer file.Close()

	return json.NewEncoder(file).Encode(c)
}

func (c *Config) decode(configFileName string) error {

	file, err := os.Open(configFileName)

	if err != nil {
		return err
	}

	defer file.Close()

	return json.NewDecoder(file).Decode(&c)
}
