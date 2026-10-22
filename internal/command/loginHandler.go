package command

import (
	"errors"
	"fmt"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("No args provided")
	}

	username := cmd.Args[0]

	s.Cfg.SetUser(username)

	fmt.Println("Username has been set")

	return nil
}