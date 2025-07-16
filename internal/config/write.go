package config

func (c *Config) write() error {

	configFileName, err := getConfigFileName()

	if err != nil {
		return err
	}

	return c.encode(configFileName)
}
