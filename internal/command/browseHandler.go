package command

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/abdullahshahad-w/blog-aggregator-go/internal/database"
)

func BrowseHandler(s *State, cmd Command, user database.User) error {
	var limit int
	if len(cmd.Args) != 0 {
		num, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return errors.New("Usage: go run . browse <optional limit>")
		}
		limit = num
	} else {
		limit = 2
	}

	posts, err := s.Db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		Name: user.Name,
		Limit: int32(limit),
	})

	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Println(post)
	}

	return nil
}