package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github/Moe-Ajam/rss-blod-aggregator/internal/database"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("args cannot be empty for the login command")
	}
	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("User has been set to: %s\n", cmd.args[0])
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("you have to pass a username to the register command")
	}

	// checking if the user already exists
	user, err := s.db.GetUser(context.Background(), sql.NullString{
		String: cmd.args[0],
	})
	if err != nil {
		log.Fatalf("something went wrong while trying to create a user: %v", err)
	}
	if user != (database.User{}) {
		fmt.Printf("User: %s already exists\n", cmd.args[0])
		os.Exit(1)
	}

	// creating the user
	createdUser, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: sql.NullString{
			String: cmd.args[0],
			Valid:  true,
		},
	})

	s.cfg.CurrentUserName = cmd.args[0]
	fmt.Printf("user %s has been created\n", createdUser.Name.String)
	fmt.Println(createdUser)

	return nil
}
