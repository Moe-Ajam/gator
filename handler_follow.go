package main

import (
	"context"
	"errors"
	"fmt"
	"github/Moe-Ajam/rss-blod-aggregator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) <= 0 {
		return errors.New("you need to pass a url to follow")
	}
	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.args[0])
	if err != nil {
		return errors.New("The requested url doesn't exist")
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("Current User: %s\n", feedFollow.UserName)
	fmt.Printf("Followed Feed: %s\n", feedFollow.FeedName)

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedsFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	for _, feedFollow := range feedFollows {
		fmt.Printf("The user %s follows %s\n", feedFollow.UserName, feedFollow.FeedName)
	}
	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) <= 0 {
		return errors.New("you need to pass a url to follow")
	}

	feedURL := cmd.args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), feedURL)
	if err != nil {
		return err
	}
	err = s.db.DeleteFollowByUser(context.Background(), database.DeleteFollowByUserParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("The feed %s has been unfollowed successfully for the user %s\n", feed.Name, user.Name)
	return nil
}
