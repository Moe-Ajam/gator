package main

import (
	"context"
	"fmt"
	"github/Moe-Ajam/rss-blod-aggregator/internal/database"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) <= 1 {
		os.Exit(1)
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
	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}
	printFeed(feed)
	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	for _, feed := range feeds {
		printFeed(feed)
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		createdBy := user.Name
		fmt.Printf("- Created By: %s\n", createdBy)
		fmt.Println("============================")
	}
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

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Fatalf("something went wrong while scraping feeds: %v\n", err)
	}
	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Fatalf("something went wrong while scraping feeds: %v\n", err)
	}
	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Fatalf("something went wrong while scraping feeds: %v\n", err)
	}

	items := rssFeed.Channel.Item
	fmt.Println("Feed Title:", rssFeed.Channel.Title)
	for _, item := range items {
		fmt.Println(item.Title)
	}
}
