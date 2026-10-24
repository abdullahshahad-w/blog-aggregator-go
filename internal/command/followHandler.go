package command

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/abdullahshahad-w/blog-aggregator-go/internal/database"
	"github.com/google/uuid"
)

func FollowHandler(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) == 0 {
		return errors.New("Usage: go run . follow <feed url>")
	}

	feedUrl := cmd.Args[0]

	feed, err := s.Db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return err
	}

	result, err := s.Db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	})

	fmt.Println(result)
	return nil
}