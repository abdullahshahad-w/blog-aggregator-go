package command

import (
	"context"
	"errors"
	"fmt"
	"time"
	"github.com/abdullahshahad-w/blog-aggregator-go/internal/database"
	"github.com/google/uuid"
)

func AddFeedHandler(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) < 2 {
		return errors.New("Usage: go run . addfeed <feed name> <feed url>")
	}

	feedName := cmd.Args[0]
	feedUrl := cmd.Args[1]

	result, err := s.Db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: feedName,
		Url: feedUrl,
		UserID: user.ID,
	})
	if err != nil {
		return err
	}

	_, err = s.Db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: result.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println("Feed was created successfully")
	fmt.Println(result)
	return nil
}