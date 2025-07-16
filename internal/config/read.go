package config

func Read() (Config, error) {

	config := Config{}

	configFileName, err := getConfigFileName()

	if err != nil {
		return config, err
	}

	err = config.decode(configFileName)

	return config, err
}
