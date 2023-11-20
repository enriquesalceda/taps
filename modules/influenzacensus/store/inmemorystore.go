package store

import "taps/domain"

type InMemoryInfluenzaStore struct {
	all map[string]domain.Census
}

func (i *InMemoryInfluenzaStore) Find(id string) bool {
	_, found := i.all[id]
	return found
}

func NewInMemoryInfluenzaStore(census map[string]domain.Census) *InMemoryInfluenzaStore {
	return &InMemoryInfluenzaStore{
		all: census,
	}
}

func (i *InMemoryInfluenzaStore) All() map[string]domain.Census {
	return i.all
}

func (i *InMemoryInfluenzaStore) Save(fieldCensus domain.Census) error {
	i.all[fieldCensus.ID] = fieldCensus
	return nil
}
