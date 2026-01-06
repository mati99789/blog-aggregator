package main

import (
	"errors"
	"fmt"
	"strings"
)

type Commands struct {
	commands map[string]CommandHandler
}

type CommandHandler func(*State, Command) error

type Command struct {
	Name string
	Args []string
}

func NewCommands() *Commands {
	return &Commands{
		commands: make(map[string]CommandHandler),
	}
}

func (c *Commands) run(s *State, cmd Command) error {
	name := strings.ToLower(cmd.Name)

	handler, ok := c.commands[name]
	if !ok {
		return errors.New(fmt.Sprintf("Unknown command: %s", name))
	}

	return handler(s, cmd)
}

func (c *Commands) register(name string, f CommandHandler) error {
	if name == "" {
		errors.New("command must have a name")
	}
	name = strings.ToLower(name)

	if _, ok := c.commands[name]; ok {
		return errors.New("command " + name + " already exists")
	}
	c.commands[name] = f

	return nil
}
