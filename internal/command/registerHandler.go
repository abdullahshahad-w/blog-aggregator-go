package command

import (
	"context"
	"errors"
	"fmt"
	"time"
	"github.com/abdullahshahad-w/blog-aggregator-go/internal/database"
	"github.com/google/uuid"
)

func RegisterHangler(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("must give arguments to register")
	}

	username := cmd.Args[0]

	result, err := s.Db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: username,
	})

	if err != nil {
		return errors.New("failed to create users, perhaps already exists")
	}

	s.Cfg.SetUser(result.Name)

	fmt.Println("user was created and set as THE user")
	
	return nil
}