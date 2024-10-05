package main

import (
	"errors"
	"fmt"
	"log"
)

type command struct {
	name string
	args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	_, ok := c.registeredCommands[name]
	if ok {
		fmt.Printf("The command %s already exists", name)
	}
	c.registeredCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.registeredCommands[cmd.name]
	if !ok {
		return errors.New("command not found")
	}

	err := f(s, cmd)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
