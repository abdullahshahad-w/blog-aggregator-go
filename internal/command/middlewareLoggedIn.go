package command

import (
	"context"
	"github.com/abdullahshahad-w/blog-aggregator-go/internal/database"
)

func MiddlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, cmd Command) error {
		user, err := s.Db.GetUserByName(context.Background(), s.Cfg.CurrentUserName)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}