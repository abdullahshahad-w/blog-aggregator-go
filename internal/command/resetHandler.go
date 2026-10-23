package command

import (
	"context"
	"fmt"
)

func ResetHandler(s *State, cmd Command) error {
	err := s.Db.ResetUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("Users table reset")
	return nil
}