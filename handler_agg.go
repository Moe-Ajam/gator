package main

import (
	"errors"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("pls add a time interval to use the command, example: 1m, 1h, 5s...")
	}
	time_between_reqs := cmd.args[0]
	fmt.Printf("Collecting feeds every %s", time_between_reqs)
	timeBetweenRequests, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return err
	}
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}
