package main

import (
	"context"
	"fmt"
	"os"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("users have been reset successfully!")
	os.Exit(0)
	return nil
}
