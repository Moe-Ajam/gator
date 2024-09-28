package main

import (
	"github/Moe-Ajam/rss-blod-aggregator/internal/config"
	"log"
	"os"
)

type state struct {
	cfg *config.Config
}

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Fatalf("Something went wrong while trying to read the config file: %v", err)
	}

	programState := state{
		cfg: &conf,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

	args := os.Args
	if len(args) <= 2 {
		log.Fatal("you need to have at least one argument inserted")
	}
	cmdName := args[1]
	cmdArgs := []string{}
	if len(args) > 2 {
		cmdArgs = args[2:]
	}
	cmd := command{
		name: cmdName,
		args: cmdArgs,
	}

	err = cmds.run(&programState, cmd)
	if err != nil {
		log.Fatalf("error while running the command %s, %v", cmd.name, err)
	}
}
