package config

func (c *Config) SetUser(name string) error {
	c.CurrentUserName = name
	return c.write()
}
