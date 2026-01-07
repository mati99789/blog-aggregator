package main

import (
	"blog_eggregator/internal/config"
	"blog_eggregator/internal/database"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()

	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	state := NewState(cfg, dbQueries)
	cmds := NewCommands()

	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: gator <command> [args...]")
		fmt.Fprintln(os.Stderr, "Available commands: login")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	args := os.Args[2:]

	cmd := Command{
		Name: cmdName,
		Args: args,
	}

	if err := cmds.run(state, cmd); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
