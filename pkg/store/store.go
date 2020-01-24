package store

import (
	"witzepedia/pkg/config"
)

type Store struct {
}

func NewStore(cfg *config.Config) *Store {
	store := &Store{}
	return store
}
