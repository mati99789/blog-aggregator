package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("login requires one argument")
	}

	err := s.Config.SetUser(cmd.Args[0])

	if err != nil {
		return err
	}

	fmt.Printf("Current user set to: %s\n", s.Config.CurrentUserName)

	return nil

}
