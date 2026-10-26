package command

import (
	"errors"
	"fmt"
	"time"
)

func AggHandler(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("Usage: go run . agg <time e.g 1s, 1m, 1h>")
	}

	timeBetweenReqs := cmd.Args[0]

	timeSt, err := time.ParseDuration(timeBetweenReqs)
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %v\n", timeSt)

	ticker := time.NewTicker(timeSt)
	for ;; <- ticker.C {
		ScrapeFeeds(s)
	}

	return nil
}