package store

import (
	"errors"
	"taps/domain"
)

type BrokenInfluenzaStore struct{}

func (b BrokenInfluenzaStore) All() map[string]domain.Census {
	return map[string]domain.Census{}
}

func (b BrokenInfluenzaStore) Save(fieldCensus domain.Census) error {
	return errors.New("something went wrong")
}

func (b BrokenInfluenzaStore) Find(id string) bool {
	return false
}

func NewBrokenInfluenzaStore() *BrokenInfluenzaStore {
	return &BrokenInfluenzaStore{}
}
