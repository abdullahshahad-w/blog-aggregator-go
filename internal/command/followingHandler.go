package command

import (
	"context"
	"fmt"
	"github.com/abdullahshahad-w/blog-aggregator-go/internal/database"
)

func FollowingHandler(s *State, cmd Command, user database.User) error {
	result, err := s.Db.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return err
	}

	fmt.Println(result)
	return nil
}