package main

import "blog_eggregator/internal/config"

type State struct {
	Config *config.Config
}

func NewState(cfg *config.Config) *State {
	return &State{Config: cfg}
}
