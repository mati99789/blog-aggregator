package main

import (
	"blog_aggregator/internal/cli"
	"blog_aggregator/internal/config"
	"blog_aggregator/internal/database"
	"blog_aggregator/internal/handlers"
	"blog_aggregator/internal/state"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	dbQueries := database.New(db)

	state := state.NewState(cfg, dbQueries)
	cmds := cli.NewCommands()

	cmds.Register("login", handlers.HandlerLogin)
	cmds.Register("register", handlers.HandlerRegister)
	cmds.Register("reset", handlers.HandlerReset)
	cmds.Register("users", handlers.HandlerListUsers)
	cmds.Register("agg", handlers.HandlerAggregate)
	cmds.Register("addfeed", handlers.HandlerAddFeed)
	cmds.Register("feeds", handlers.HandlerListFeeds)

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: gator <command> [args...]")
		fmt.Fprintln(os.Stderr, "Available commands: login")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	args := os.Args[2:]

	cmd := cli.Command{
		Name: cmdName,
		Args: args,
	}

	if err := cmds.Run(state, cmd); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
