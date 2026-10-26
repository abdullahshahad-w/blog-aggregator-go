package command

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/abdullahshahad-w/blog-aggregator-go/internal/database"
	"github.com/abdullahshahad-w/blog-aggregator-go/internal/rss"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func ScrapeFeeds(s *State) error {
	nextFeed, err := s.Db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	err = s.Db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{
			Time: time.Now(),
			Valid: true,
		},
		UpdatedAt: time.Now(),
		ID: nextFeed.ID,
	})
	if err != nil {
		return err
	}

	rssFeed, err := rss.FetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}

	for _, item := range rssFeed.Channel.Item {
		post, err := s.Db.CreatePost(context.Background(), database.CreatePostParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Title: item.Title,
		Url: item.Link,
		Description: sql.NullString{
			String: item.Link,
			Valid: item.Link != "",
		},
		PublishedAt: sql.NullTime{
			Time: time.Now(),
			Valid: false,
		},
		FeedID: nextFeed.ID,
		})

		if err != nil {
			var pqErr *pq.Error
			if errors.As(err, &pqErr) && pqErr.Code == "23505" {
				continue
			}
			return err
		}

		fmt.Println("Created Post.")
		fmt.Println(post)
	}

	return nil
}