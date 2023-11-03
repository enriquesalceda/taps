package influenzacensus

import "errors"

type FieldCensusParameters struct {
	ID            string
	FirstLastName string
	LastLastName  string
	FirstName     string
	DOB           string
	State         string
	Gender        string
	Number        int
}

type FieldCensus struct {
	ID            string
	FirstLastName string
	LastLastName  string
	FirstName     string
	DOB           string
	State         string
	Gender        string
	Number        int
}

type InfluenzaCensusTaker struct {
	store CensusStore
}

func NewInfluenzaCensusTaker(store CensusStore) *InfluenzaCensusTaker {
	return &InfluenzaCensusTaker{store: store}
}

func (t *InfluenzaCensusTaker) Take(
	fieldCensusParameters *FieldCensusParameters,
) error {
	fieldCensus := &FieldCensus{
		ID:            fieldCensusParameters.ID,
		FirstLastName: fieldCensusParameters.FirstLastName,
		LastLastName:  fieldCensusParameters.LastLastName,
		FirstName:     fieldCensusParameters.FirstName,
		DOB:           fieldCensusParameters.DOB,
		State:         fieldCensusParameters.State,
		Gender:        fieldCensusParameters.Gender,
		Number:        fieldCensusParameters.Number,
	}

	if fieldCensus.ID == "" {
		return errors.New("No ID")
	}

	if fieldCensus.FirstLastName == "" {
		return errors.New("No FirstLastName")
	}

	if fieldCensus.LastLastName == "" {
		return errors.New("No LastLastName")
	}

	if fieldCensus.FirstName == "" {
		return errors.New("No FirstName")
	}

	if fieldCensus.DOB == "" {
		return errors.New("No DOB")
	}

	if fieldCensus.Gender == "" {
		return errors.New("No Gender")
	}

	if fieldCensus.State == "" {
		return errors.New("No State")
	}

	if fieldCensus.Number == 0 {
		return errors.New("No Number")
	}

	return t.store.Save(fieldCensus)
}

// -- store
type InMemoryInfluenzaStore struct {
	all map[string]InfluenzaCensus
}

type InfluenzaCensus struct {
	ID            string
	FirstLastName string
	LastLastName  string
	FirstName     string
	DOB           string
	State         string
	Gender        string
	Number        int
}

type CensusStore interface {
	All() []InfluenzaCensus
	Save(fieldCensus *FieldCensus) error
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

func (i *InMemoryInfluenzaStore) Save(fieldCensus *FieldCensus) error {
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
