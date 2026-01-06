package main

import (
	"blog_eggregator/internal/config"
	"fmt"
	"os"
)

func main() {
	cfg, err := config.Read()

	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	state := NewState(cfg)
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
