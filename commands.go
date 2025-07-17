package main

type command struct {
	name      string
	arguments []string
}

type commands struct {
	availableCommands map[string]func(s *state, cmd command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.availableCommands[cmd.name]

	if !ok {
		return nil
	}

	return handler(s, cmd)
}

func (c *commands) register(name string, handler func(s *state, cmd command) error) {
	c.availableCommands[name] = handler
}
