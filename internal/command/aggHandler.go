package command

import (
	"context"
	"fmt"

	"github.com/abdullahshahad-w/blog-aggregator-go/internal/rss"
)

func AggHandler(s *State, cmd Command) error {

	url := "https://www.wagslane.dev/index.xml"

	rssStruct, err := rss.FetchFeed(context.Background(), url)
	if err != nil {
		return err
	}

	fmt.Println(*rssStruct)
	return nil
}