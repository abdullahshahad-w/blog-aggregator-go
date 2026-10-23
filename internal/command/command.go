package command

import "errors"

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	CmdMap map[string]func (*State, Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	if val, ok := c.CmdMap[cmd.Name]; !ok {
		return errors.New("This command is not available")
	} else {
		return val(s, cmd)
	}
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.CmdMap[name] = f
}