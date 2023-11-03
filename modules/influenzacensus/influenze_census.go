package influenzacensus

import (
	"errors"
)

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

	err := ValidatePresence(fieldCensus)
	if err != nil {
		return err
	}

	return t.store.Save(fieldCensus)
}

func ValidatePresence(fieldCensus *FieldCensus) error {
	validationErrors := []error{}

	if fieldCensus.ID == "" {
		validationErrors = append(validationErrors, errors.New("No ID"))
	}

	if fieldCensus.FirstLastName == "" {
		validationErrors = append(validationErrors, errors.New("No FirstLastName"))
	}

	if fieldCensus.LastLastName == "" {
		validationErrors = append(validationErrors, errors.New("No LastLastName"))
	}

	if fieldCensus.FirstName == "" {
		validationErrors = append(validationErrors, errors.New("No FirstName"))
	}

	if fieldCensus.DOB == "" {
		validationErrors = append(validationErrors, errors.New("No DOB"))
	}

	if fieldCensus.Gender == "" {
		validationErrors = append(validationErrors, errors.New("No Gender"))
	}

	if fieldCensus.State == "" {
		validationErrors = append(validationErrors, errors.New("No State"))
	}

	if fieldCensus.Number == 0 {
		validationErrors = append(validationErrors, errors.New("No Number"))
	}

	if len(validationErrors) > 0 {
		return errors.Join(validationErrors...)
	}

	return nil
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
