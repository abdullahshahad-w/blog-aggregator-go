package command

import (
	"context"
	"fmt"
)

func UsersHandler(s *State, cmd Command) error {
	result, err := s.Db.GetAllUsers(context.Background())
	if err != nil {
		return err
	}

	currentUser := s.Cfg.CurrentUserName

	for _, user := range result {
		if user.Name == currentUser {
			fmt.Printf("* %s (current)\n", user.Name)
		}
		fmt.Printf("* %s\n", user.Name)
	}

	return nil
}