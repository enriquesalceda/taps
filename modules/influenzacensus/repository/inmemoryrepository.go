package repository

import (
	"taps/domain"
	"taps/utils/clk"
)

type InMemoryPrimaryKey struct {
	CurpID string
	Date   string
}

type InMemoryInfluenzaStore struct {
	all   map[InMemoryPrimaryKey]domain.Census
	clock clk.Clock
}

func NewInMemoryInfluenzaStore(census map[InMemoryPrimaryKey]domain.Census, clock clk.Clock) *InMemoryInfluenzaStore {
	return &InMemoryInfluenzaStore{
		all:   census,
		clock: clock,
	}
}

func (i *InMemoryInfluenzaStore) Find(id string, date string) (domain.Census, bool, error) {
	item, found := i.all[NewInMemoryPrimaryKey(id, date)]
	if !found {
		return domain.Census{}, found, nil
	}

	return item, found, nil
}

func (i *InMemoryInfluenzaStore) Save(fieldCensus domain.Census) error {
	i.all[NewInMemoryPrimaryKey(fieldCensus.ID, i.clock.Now().Format("2006-01-02"))] = fieldCensus
	return nil
}

func NewInMemoryPrimaryKey(id string, date string) InMemoryPrimaryKey {
	return InMemoryPrimaryKey{
		CurpID: id,
		Date:   date,
	}
}
