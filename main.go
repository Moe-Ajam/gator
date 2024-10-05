package main

import (
	"database/sql"
	"github/Moe-Ajam/rss-blod-aggregator/internal/config"
	"github/Moe-Ajam/rss-blod-aggregator/internal/database"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Fatalf("something went wrong while trying to read the config file: %v", err)
	}

	db, err := sql.Open("postgres", conf.DBURL)
	if err != nil {
		log.Fatalf("something went wrong while creating a db connection: %v", err)
	}
	dbQueries := database.New(db)

	programState := state{
		db:  dbQueries,
		cfg: &conf,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)

	args := os.Args
	if len(args) <= 1 {
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
