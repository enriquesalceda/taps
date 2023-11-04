package store

import (
	"errors"
	"taps/domain"
)

type BrokenInfluenzaStore struct{}

func (b BrokenInfluenzaStore) All() []InfluenzaCensus {
	return []InfluenzaCensus{}
}

func (b BrokenInfluenzaStore) Save(fieldCensus domain.FieldCensus) error {
	return errors.New("something went wrong")
}

func (b BrokenInfluenzaStore) Find(id string) bool {
	return false
}

func NewBrokenInfluenzaStore() *BrokenInfluenzaStore {
	return &BrokenInfluenzaStore{}
}