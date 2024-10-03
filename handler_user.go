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

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("args cannot be empty for the login command")
	}

	// checking if the user exists in the database
	user, _ := s.db.GetUser(context.Background(), cmd.args[0])
	fmt.Println("name:", user.Name, "id:", user.ID, "updatedAt:", user.UpdatedAt)
	if user == (database.User{}) {
		fmt.Printf("user %s doesn't exist\n", cmd.args[0])
		os.Exit(1)
	}

	// setting the current user to the user passed in args
	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("user has been set to: %s\n", cmd.args[0])
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("you have to pass a username to the register command")
	}

	// checking if the user already exists
	user, _ := s.db.GetUser(context.Background(), cmd.args[0])
	if user != (database.User{}) {
		fmt.Printf("user: %s already exists\n", cmd.args[0])
		os.Exit(1)
	}

	// creating the user
	createdUser, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.args[0],
	})
	if err != nil {
		// log.Fatalf("something went wrong while trying to register a user: %v", err)
		return err
	}

	s.cfg.SetUser(createdUser.Name)
	fmt.Printf("user %s has been created\n", createdUser.Name)
	fmt.Println(createdUser)

	return nil
}
