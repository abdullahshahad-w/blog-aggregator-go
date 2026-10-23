package command

import (
	"context"
	"errors"
	"fmt"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("No args provided")
	}

	username := cmd.Args[0]

	result, err := s.Db.GetUserByName(context.Background(), username)

	if err != nil {
		return errors.New("user not found in the db")
	}

	s.Cfg.SetUser(result.Name)

	fmt.Println("Username has been set")

	return nil
}