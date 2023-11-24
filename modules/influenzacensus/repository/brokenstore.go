package repository

import (
	"errors"
	"taps/domain"
)

type BrokenInfluenzaStore struct {
	when string
}

func NewBrokenInfluenzaStore(when string) *BrokenInfluenzaStore {
	if when != "save" && when != "find" {
		panic("invalid when for broken store")
	}
	return &BrokenInfluenzaStore{when: when}
}

func (b BrokenInfluenzaStore) Save(fieldCensus domain.Census) error {
	if b.when == "save" {
		return errors.New("something went wrong")
	}

	return nil
}

func (b BrokenInfluenzaStore) Find(id string, date string) (domain.Census, bool, error) {
	if b.when == "find" {
		return domain.Census{}, false, errors.New("something went wrong")
	}

	return domain.Census{}, false, nil
}
