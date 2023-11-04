package store

import "taps/domain"

type InMemoryInfluenzaStore struct {
	all map[string]InfluenzaCensus
}

func NewInMemoryInfluenzaStore() *InMemoryInfluenzaStore {
	return &InMemoryInfluenzaStore{
		all: map[string]InfluenzaCensus{},
	}
}

func (i *InMemoryInfluenzaStore) All() []InfluenzaCensus {
	var allCensus []InfluenzaCensus
	for _, census := range i.all {
		allCensus = append(allCensus, census)
	}
	return allCensus
}

func (i *InMemoryInfluenzaStore) Save(fieldCensus domain.FieldCensus) error {
	influenzaCensus := InfluenzaCensus{
		ID:            fieldCensus.ID,
		FirstLastName: fieldCensus.FirstLastName,
		LastLastName:  fieldCensus.LastLastName,
		FirstName:     fieldCensus.FirstName,
		DOB:           fieldCensus.DOB,
		State:         fieldCensus.State,
		Gender:        fieldCensus.Gender,
		Number:        fieldCensus.Number,
	}

	i.all[fieldCensus.ID] = influenzaCensus
	return nil
}
