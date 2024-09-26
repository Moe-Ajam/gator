package main

import (
	"errors"
	"fmt"
	"github/Moe-Ajam/rss-blod-aggregator/internal/config"
	"log"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Fatalf("Something went wrong while trying to read the config file: %v", err)
	}
	fmt.Printf("Read config file: %v\n", conf)
	err = conf.SetUser("moe")
	if err != nil {
		log.Fatalf("Something went wrong while setting the user: %v", err)
	}
	conf1, err := config.Read()
	fmt.Printf("Read config file again: %v\n", conf1)
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Args cannot be empty for the login command")
	}
	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("User has been set to: %s", cmd.args[0])
	return nil
}

// registers a new command
func (c *commands) register(name string, f func(*state, command) error) {
	_, ok := c.cmds[name]
	if ok {
		fmt.Printf("The command %s already exists", name)
	}
	c.cmds[name] = f
}

// runs the provided command with the provided state if it exists
func (c *commands) run(s *state, cmd command) error {
	return nil
}
