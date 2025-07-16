package config

import (
	"encoding/json"
	"os"
)

func Read() (Config, error) {

	config := Config{}

	configFileName, err := getConfigFileName()

	if err != nil {
		return config, err
	}

	file, err := os.Open(configFileName)

	if err != nil {
		return config, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&config)

	if err != nil {
		return config, err
	}

	return config, err
}
