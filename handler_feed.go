package main

import (
	"context"
	"errors"
	"fmt"
	"github/Moe-Ajam/rss-blod-aggregator/internal/database"
	"os"
	"time"

	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) <= 1 {
		os.Exit(1)
		return errors.New("you have to pass a name and a url to run the addFeed command")
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		fmt.Println(err)
		return err
	}

	name := cmd.args[0]
	url := cmd.args[1]
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}
	printFeed(feed)
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("- ID: %v\n", feed.ID)
	fmt.Printf("- CreatedAt: %v\n", feed.CreatedAt)
	fmt.Printf("- UpdatedAt: %v\n", feed.UpdatedAt)
	fmt.Printf("- Name: %v\n", feed.Name)
	fmt.Printf("- Url: %v\n", feed.Url)
	fmt.Printf("- User ID: %v\n", feed.UserID)
}
