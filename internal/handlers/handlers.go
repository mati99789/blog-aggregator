package handlers

import (
	"blog_aggregator/external/api"
	"blog_aggregator/internal/cli"
	"blog_aggregator/internal/database"
	"blog_aggregator/internal/state"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func HandlerLogin(s *state.State, cmd cli.Command) error {
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

func HandlerRegister(s *state.State, cmd cli.Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("register requires one argument")
	}

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

func HandlerReset(s *state.State, cmd cli.Command) error {
	err := s.Db.DeleteAllUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("All users deleted")

	return nil

}

func HandlerListUsers(s *state.State, cmd cli.Command) error {
	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	CurrentUserName := s.Config.CurrentUserName

	for _, user := range users {
		if user.Name == CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %s\n", user.Name)
	}

	return nil
}

func HandlerAggregate(s *state.State, cmd cli.Command) error {
	ressFeed, err := api.FetchRSSFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Println(ressFeed)

	return nil
}
