package main

import (
	"blog_eggregator/internal/config"
	"blog_eggregator/internal/database"
)

type State struct {
	Db     *database.Queries
	Config *config.Config
}

func NewState(cfg *config.Config, db *database.Queries) *State {
	return &State{Config: cfg, Db: db}
}
