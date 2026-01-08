package state

import (
	"blog_aggregator/internal/config"
	"blog_aggregator/internal/database"
)

type State struct {
	Db     *database.Queries
	Config *config.Config
}

func NewState(cfg *config.Config, db *database.Queries) *State {
	return &State{Config: cfg, Db: db}
}
