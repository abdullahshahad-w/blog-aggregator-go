package command

import (
	"context"
	"fmt"
)

func FeedsHandler(s *State, cmd Command) error {
	result , err := s.Db.GetAllFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range result {
		fmt.Println(feed)
	}

	return nil
}