package middleware

import (
	"blog_aggregator/internal/cli"
	"blog_aggregator/internal/database"
	"blog_aggregator/internal/state"
	"context"
)

func MiddlewareLoggedIn(handler func(s *state.State, cmd cli.Command, user database.User) error) func(*state.State, cli.Command) error {
	return func(s *state.State, cmd cli.Command) error {
		user, err := s.Db.GetUser(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
