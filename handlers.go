package main

import (
	"blog_eggregator/internal/database"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("login requires one argument")
	}

	_, err := s.Db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return errors.New("user does not exist")
	}

	err = s.Config.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Current user set to: %s\n", s.Config.CurrentUserName)

	return nil
}

func handlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("register requires one argument")
	}

	// Sprawdź czy użytkownik już istnieje
	_, err := s.Db.GetUser(context.Background(), cmd.Args[0])
	if err == nil {
		return errors.New("user already exists")
	}

	user, err := s.Db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	})

	if err != nil {
		return err
	}

	err = s.Config.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("User created: %s\n", user.Name)

	return nil
}
