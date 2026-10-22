package command

import "errors"

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	cmdMap map[string]func (*State, Command) error
}

func (c *Commands) run(s *State, cmd Command) error {
	if val, ok := c.cmdMap[cmd.Name]; !ok {
		return errors.New("This command is not available")
	} else {
		val(s, cmd)
		return nil
	}
}

func (c *Commands) register(name string, f func(*State, Command) error) {
	c.cmdMap[name] = f
}