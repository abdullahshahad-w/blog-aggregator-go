package command

import (
	"context"
	"errors"
	"fmt"
	"github.com/abdullahshahad-w/blog-aggregator-go/internal/database"
)

func UnfollowHandler(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) == 0 {
		return errors.New("Usage: go run . unfollow <feed url>")
	}

	feedUrl := cmd.Args[0]

	feed, err := s.Db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return err
	}

	if err := s.Db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}); err != nil {
		return err
	}

	fmt.Println("Deleted Successfully")
	return nil
}